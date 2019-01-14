package staticUtil

import (
	"github.com/ingmardrewing/staticIntf"
	log "github.com/sirupsen/logrus"
)

func NewPagesContainerCollectionTool(
	pcc staticIntf.PagesContainerCollection) *pagesContainerCollectionTool {
	pt := new(pagesContainerCollectionTool)
	pt.collection = pcc
	return pt
}

type pagesContainerCollectionTool struct {
	collection staticIntf.PagesContainerCollection
}

func (c *pagesContainerCollectionTool) ContainersOrderedByVariants(variants ...string) []staticIntf.PagesContainer {
	log.Debug("ContainersOrderedByVariants, nr of containers:", len(c.collection.Containers()))
	orderedContainers := []staticIntf.PagesContainer{}
	for _, v := range variants {
		log.Debug("ContainersOrderedByVariants - looping through variant:", v)
		container := c.getContainersByVariant(v)
		if container != nil {
			orderedContainers = append(orderedContainers, container...)
		}
	}
	return orderedContainers
}

func (c *pagesContainerCollectionTool) getContainersByVariant(v string) []staticIntf.PagesContainer {
	containers := []staticIntf.PagesContainer{}
	for _, co := range c.collection.Containers() {
		if co.Variant() == v {
			containers = append(containers, co)
		}
	}
	return containers
}

func (c *pagesContainerCollectionTool) getContainerByVariant(v string) staticIntf.PagesContainer {
	for _, co := range c.collection.Containers() {
		log.Debug("getContainerByVariant: ", co.Variant(), "==?", v)
		if co.Variant() == v {
			log.Debug("getContainerByVariant, returning: ", co.Variant())
			return co
		}
	}
	return nil
}

func (c *pagesContainerCollectionTool) GetNaviPagesByVariant(v string) []staticIntf.Page {
	nps := []staticIntf.Page{}
	for _, c := range c.getContainersByVariant(v) {
		nps = append(nps, c.NaviPages()...)
	}
	return nps
}

func (c *pagesContainerCollectionTool) GetPagesByVariant(v string) []staticIntf.Page {
	ps := []staticIntf.Page{}
	for _, c := range c.getContainersByVariant(v) {
		ps = append(ps, c.Pages()...)
	}
	return ps
}
