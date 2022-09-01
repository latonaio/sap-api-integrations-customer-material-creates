package sap_api_input_reader

import (
	"sap-api-integrations-customer-material-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToCustomerMaterial() *requests.CustomerMaterial {
	data := sdc.CustomerMaterial
	return &requests.CustomerMaterial{
			SalesOrganization:              data.SalesOrganization,
			DistributionChannel:            data.DistributionChannel,
			Customer:                       data.Customer,
			Material:                       data.Material,
			MaterialByCustomer:             data.MaterialByCustomer,
			MaterialDescriptionByCustomer:  data.MaterialDescriptionByCustomer,
			Plant:                          data.Plant,
			DeliveryPriority:               data.DeliveryPriority,
			MinDeliveryQtyInBaseUnit:       data.MinDeliveryQtyInBaseUnit,
			BaseUnit:                       data.BaseUnit,
			PartialDeliveryIsAllowed:       data.PartialDeliveryIsAllowed,
			MaxNmbrOfPartialDelivery:       data.MaxNmbrOfPartialDelivery,
			UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
			OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
			UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
			CustomerMaterialItemUsage:      data.CustomerMaterialItemUsage,
			SalesUnit:                      data.SalesUnit,
			SalesQtyToBaseQtyDnmntr:        data.SalesQtyToBaseQtyDnmntr,
			SalesQtyToBaseQtyNmrtr:         data.SalesQtyToBaseQtyNmrtr,
	}
}
