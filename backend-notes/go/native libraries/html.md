# Description

Ce package permet de parser un fichier HTML.

# Syntaxe

```go
// io.Reader : File or web request (giving back an html)
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var link Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link.Href = attr.Val
		}
	}
	link.Text = extractText(n)
	return link
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode { // Nous ne sommes pas intéressés par les commentaires, ...
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += extractText(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...) // linkNodes retourne un slice donc pour l'append il faut une valeur simple c'est pourquoi on utilise le spread operator ...
	}
	return ret
}
```
