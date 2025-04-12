package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

func ToResourceItemProto(resource do.Resource) *common.ResourceItem {
	if resource == nil {
		return nil
	}
	return &common.ResourceItem{
		Id:        resource.GetID(),
		Name:      resource.GetName(),
		Path:      resource.GetPath(),
		Status:    common.GlobalStatus(resource.GetStatus()),
		Remark:    resource.GetRemark(),
		CreatedAt: resource.GetCreatedAt().Format(time.DateTime),
		UpdatedAt: resource.GetUpdatedAt().Format(time.DateTime),
	}
}

func ToResourceItemProtoList(resources []do.Resource) []*common.ResourceItem {
	result := make([]*common.ResourceItem, 0, len(resources))
	for _, resource := range resources {
		result = append(result, ToResourceItemProto(resource))
	}
	return result
}

func ToMenuTreeProto(menus []do.Menu) []*common.MenuTreeItem {
	menuMap := make(map[uint32]do.Menu)
	for _, menu := range menus {
		menuMap[menu.GetID()] = menu
	}

	// 构建树形结构
	var roots []*common.MenuTreeItem
	for _, menu := range menus {
		if menu.GetParentID() == 0 {
			roots = append(roots, convertMenuToTreeItemWithMap(menu, menuMap))
		}
	}

	return roots
}

func convertMenuToTreeItemWithMap(menu do.Menu, menuMap map[uint32]do.Menu) *common.MenuTreeItem {
	treeItem := &common.MenuTreeItem{
		Id:       menu.GetID(),
		Name:     menu.GetName(),
		Path:     menu.GetPath(),
		Status:   common.GlobalStatus(menu.GetStatus()),
		Icon:     menu.GetIcon(),
		Children: nil,
	}

	// 查找所有子菜单
	for _, m := range menuMap {
		if m.GetParentID() == menu.GetID() {
			if treeItem.Children == nil {
				treeItem.Children = make([]*common.MenuTreeItem, 0)
			}
			treeItem.Children = append(treeItem.Children, convertMenuToTreeItemWithMap(m, menuMap))
		}
	}

	return treeItem
}
