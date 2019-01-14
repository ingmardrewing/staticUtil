package staticUtil

import "github.com/ingmardrewing/staticIntf"

func NewPagesContainerTool(container staticIntf.PagesContainer) staticIntf.PagesContainerTool {
	t := new(pagesContainerTool)
	t.container = container
	return t
}

type pagesContainerTool struct {
	container staticIntf.PagesContainer
}

func (c *pagesContainerTool) GetIndexOfPage(p staticIntf.Page) int {
	return c.getIndex(c.container.Pages(), p)
}

func (c *pagesContainerTool) GetIndexOfNaviPage(p staticIntf.Page) int {
	return c.getIndex(c.container.NaviPages(), p)
}

func (c *pagesContainerTool) getIndex(
	pgs []staticIntf.Page,
	p staticIntf.Page) int {

	for i, l := range pgs {
		if l.Link() == p.Link() {
			return i
		}
	}
	return -1
}

func (c *pagesContainerTool) GetLastPage(p staticIntf.Page) staticIntf.Page {
	if c.GetIndexOfPage(p) != -1 {
		return c.getLastOf(c.container.Pages())
	} else if c.GetIndexOfNaviPage(p) != -1 {
		return c.getLastOf(c.container.NaviPages())
	}
	return nil
}

func (c *pagesContainerTool) getLastOf(pgs []staticIntf.Page) staticIntf.Page {
	if len(pgs) > 0 {
		return pgs[len(pgs)-1]
	}
	return nil
}

func (c *pagesContainerTool) GetFirstPage(p staticIntf.Page) staticIntf.Page {
	if c.GetIndexOfPage(p) != -1 {
		return c.getFirstOf(c.container.Pages())
	} else if c.GetIndexOfNaviPage(p) != -1 {
		return c.getFirstOf(c.container.NaviPages())
	}
	return nil
}

func (c *pagesContainerTool) getFirstOf(pgs []staticIntf.Page) staticIntf.Page {

	if len(pgs) > 0 {
		return pgs[0]
	}
	return nil
}

func (c *pagesContainerTool) GetPageBefore(p staticIntf.Page) staticIntf.Page {
	if c.GetIndexOfPage(p) != -1 {
		return c.pageBefore(c.container.Pages(), p)
	} else if c.GetIndexOfNaviPage(p) != -1 {
		return c.pageBefore(c.container.NaviPages(), p)
	}
	return nil
}

func (c *pagesContainerTool) pageBefore(pgs []staticIntf.Page, p staticIntf.Page) staticIntf.Page {
	i := c.getIndex(pgs, p)
	if i > 0 {
		return pgs[i-1]
	}
	return nil
}

func (c *pagesContainerTool) GetPageAfter(p staticIntf.Page) staticIntf.Page {
	if c.GetIndexOfPage(p) != -1 {
		return c.pageAfter(c.container.Pages(), p)
	} else if c.GetIndexOfNaviPage(p) != -1 {
		return c.pageAfter(c.container.NaviPages(), p)
	}
	return nil
}

func (c *pagesContainerTool) pageAfter(pgs []staticIntf.Page, p staticIntf.Page) staticIntf.Page {
	i := c.getIndex(pgs, p)
	if i < len(pgs)-1 {
		return pgs[i+1]
	}
	return nil
}

func (c *pagesContainerTool) SiblingPages(p staticIntf.Page) []staticIntf.Page {
	if c.GetIndexOfPage(p) != -1 {
		return c.container.Pages()
	} else if c.GetIndexOfNaviPage(p) != -1 {
		return c.container.NaviPages()
	}
	return nil
}
