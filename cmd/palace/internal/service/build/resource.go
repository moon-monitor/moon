package build

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

func ToResourceItemProto(resource *system.Resource) *common.ResourceItem {
	if resource == nil {
		return nil
	}
	return &common.ResourceItem{
		Id:        resource.ID,
		Name:      resource.Name,
		Path:      resource.Path,
		Status:    common.GlobalStatus(resource.Status),
		Remark:    resource.Remark,
		CreatedAt: resource.CreatedAt.Format(time.DateTime),
		UpdatedAt: resource.UpdatedAt.Format(time.DateTime),
	}
}

func ToResourceItemProtoList(resources []*system.Resource) []*common.ResourceItem {
	result := make([]*common.ResourceItem, 0, len(resources))
	for _, resource := range resources {
		result = append(result, ToResourceItemProto(resource))
	}
	return result
}

func ToMenuTreeProto(menus []*system.Menu) []*common.MenuTreeItem {
	menuMap := make(map[uint32]*system.Menu)
	for _, menu := range menus {
		menuMap[menu.ID] = menu
	}

	// 构建树形结构
	var roots []*common.MenuTreeItem
	for _, menu := range menus {
		if menu.ParentID == 0 {
			roots = append(roots, convertMenuToTreeItemWithMap(menu, menuMap))
		}
	}

	return roots
}

func convertMenuToTreeItemWithMap(menu *system.Menu, menuMap map[uint32]*system.Menu) *common.MenuTreeItem {
	treeItem := &common.MenuTreeItem{
		Id:       menu.ID,
		Name:     menu.Name,
		Path:     menu.Path,
		Status:   common.GlobalStatus(menu.Status),
		Icon:     menu.Icon,
		Children: nil,
	}

	// 查找所有子菜单
	for _, m := range menuMap {
		if m.ParentID == menu.ID {
			if treeItem.Children == nil {
				treeItem.Children = make([]*common.MenuTreeItem, 0)
			}
			treeItem.Children = append(treeItem.Children, convertMenuToTreeItemWithMap(m, menuMap))
		}
	}

	return treeItem
}
