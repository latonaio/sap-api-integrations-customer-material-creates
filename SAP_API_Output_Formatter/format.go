package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-customer-material-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToCustomerMaterial(raw []byte, l *logger.Logger) (*CustomerMaterial, error) {
	pm := &responses.CustomerMaterial{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Bank. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D

	customerMaterial := &CustomerMaterial{
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

	return customerMaterial, nil
}
