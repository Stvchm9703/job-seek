package job_classification

import (
	"fmt"

	"github.com/samber/lo"
)

func FindParentCategory(input string) string {

	originalCategories, hasCat := lo.Find(JobClassList, func(itm JobClassificationSet) bool {
		return itm.Label == input || fmt.Sprintf("%d", itm.Id) == input
	})

	if !hasCat {
		return ""
	}

	parent, hasParent := lo.Find(JobClassList, func(itm JobClassificationSet) bool {
		return itm.Label == originalCategories.Hierarchy
	})
	if !hasParent {
		return fmt.Sprintf("%d", originalCategories.Id)
	}
	return fmt.Sprintf("%d", parent.Id)
}
