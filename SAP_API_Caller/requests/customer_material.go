package requests

type CustomerMaterial struct {
			SalesOrganization              string `json:"SalesOrganization"`
			DistributionChannel            string `json:"DistributionChannel"`
			Customer                       string `json:"Customer"`
			Material                       string `json:"Material"`
			MaterialByCustomer             *string `json:"MaterialByCustomer"`
			MaterialDescriptionByCustomer  *string `json:"MaterialDescriptionByCustomer"`
			Plant                          *string `json:"Plant"`
			DeliveryPriority               *string `json:"DeliveryPriority"`
			MinDeliveryQtyInBaseUnit       *string `json:"MinDeliveryQtyInBaseUnit"`
			BaseUnit                       *string `json:"BaseUnit"`
			PartialDeliveryIsAllowed       *string `json:"PartialDeliveryIsAllowed"`
			MaxNmbrOfPartialDelivery       *string `json:"MaxNmbrOfPartialDelivery"`
			UnderdelivTolrtdLmtRatioInPct  *string `json:"UnderdelivTolrtdLmtRatioInPct"`
			OverdelivTolrtdLmtRatioInPct   *string `json:"OverdelivTolrtdLmtRatioInPct"`
			UnlimitedOverdeliveryIsAllowed *bool   `json:"UnlimitedOverdeliveryIsAllowed"`
			CustomerMaterialItemUsage      *string `json:"CustomerMaterialItemUsage"`
			SalesUnit                      *string `json:"SalesUnit"`
			SalesQtyToBaseQtyDnmntr        *string `json:"SalesQtyToBaseQtyDnmntr"`
			SalesQtyToBaseQtyNmrtr         *string `json:"SalesQtyToBaseQtyNmrtr"`
}
