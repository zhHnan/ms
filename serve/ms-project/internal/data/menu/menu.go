package menu

import "github.com/jinzhu/copier"

type ProjectMenu struct {
	Id         int64
	Pid        int64
	Title      string
	Icon       string
	Url        string
	FilePath   string
	Params     string
	Node       string
	Sort       int
	Status     int
	CreateBy   int64
	IsInner    int
	Values     string
	ShowSlider int
}

func (*ProjectMenu) TableName() string {
	return "ms_project_menu"
}

type ProjectMenuChild struct {
	ProjectMenu
	Children []*ProjectMenuChild
}

func CovertChild(pms []*ProjectMenu) []*ProjectMenuChild {
	var pmcs []*ProjectMenuChild
	copier.Copy(&pmcs, pms)
	var childPmcs []*ProjectMenuChild
	//递归
	for _, v := range pmcs {
		if v.Pid == 0 {
			pmc := &ProjectMenuChild{}
			copier.Copy(pmc, v)
			childPmcs = append(childPmcs, pmc)
		}
	}
	toChild(childPmcs, pmcs)
	return childPmcs
}

func toChild(childPmcs []*ProjectMenuChild, pmcs []*ProjectMenuChild) {
	for _, pmc := range childPmcs {
		for _, pm := range pmcs {
			if pmc.Id == pm.Pid {
				child := &ProjectMenuChild{}
				copier.Copy(child, pm)
				pmc.Children = append(pmc.Children, child)
			}
		}
		toChild(pmc.Children, pmcs)
	}
}
