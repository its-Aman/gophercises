package HTMLparser

import (
	"io"
	"log"
	"strings"

	xhtml "golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := xhtml.Parse(r)

	if err != nil {
		log.Fatal("Error while parsing the file", err)
		panic(err)
	}

	nodes := linkNodes(doc)
	var links []Link

	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	// fmt.Printf("%+v", links)

	return links, nil
}

func linkNodes(n *xhtml.Node) []*xhtml.Node {
	if n.Type == xhtml.ElementNode && n.Data == "a" {
		return []*xhtml.Node{n}
	}

	var ret []*xhtml.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}

func buildLink(n *xhtml.Node) Link {
	var ret Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}

	}

	ret.Text = buildText(n)
	return ret
}

func buildText(n *xhtml.Node) string {
	if n.Type == xhtml.TextNode {
		return n.Data
	}

	if n.Type != xhtml.ElementNode {
		return ""
	}

	var ret string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += buildText(c) + " "
	}

	return strings.Join(strings.Fields(ret), " ")
}
