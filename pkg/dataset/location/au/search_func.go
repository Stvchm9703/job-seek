package au

import (
	"fmt"

	"github.com/samber/lo"
)

func FindParentLocation(input string) string {
	original, hasCat := lo.Find(PostcodeList, func(itm Postcode) bool {
		return itm.PlaceName == input || fmt.Sprintf("%d", itm.Postcode) == input
	})

	if !hasCat {
		return ""
	}

	if original.IsCapital {
		return fmt.Sprintf("%s %s %d", original.PlaceName, original.StateCode, original.Postcode)
	}

	nearCap, hasNearCap := lo.Find(PostcodeList, func(itm Postcode) bool {
		return itm.StateCode == original.StateCode && itm.IsCapital == true
	})
	if !hasNearCap {
		return fmt.Sprintf("%s %s %d", original.PlaceName, original.StateCode, original.Postcode)
	}
	return fmt.Sprintf("%s %s %d", nearCap.PlaceName, nearCap.StateCode, nearCap.Postcode)

}
