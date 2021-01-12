package sentinel_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/securityinsight/mgmt/2019-01-01-preview/securityinsight"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/sentinel/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type SentinelAlertRuleMsSecurityIncidentResource struct {
}

func TestAccSentinelAlertRuleMsSecurityIncident_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sentinel_alert_rule_ms_security_incident", "test")
	r := SentinelAlertRuleMsSecurityIncidentResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSentinelAlertRuleMsSecurityIncident_complete(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sentinel_alert_rule_ms_security_incident", "test")
	r := SentinelAlertRuleMsSecurityIncidentResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSentinelAlertRuleMsSecurityIncident_update(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sentinel_alert_rule_ms_security_incident", "test")
	r := SentinelAlertRuleMsSecurityIncidentResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.complete(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSentinelAlertRuleMsSecurityIncident_requiresImport(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sentinel_alert_rule_ms_security_incident", "test")
	r := SentinelAlertRuleMsSecurityIncidentResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.RequiresImportErrorStep(r.requiresImport),
	})
}

func TestAccSentinelAlertRuleMsSecurityIncident_withAlertRuleTemplateGuid(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sentinel_alert_rule_ms_security_incident", "test")
	r := SentinelAlertRuleMsSecurityIncidentResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.alertRuleTemplateGuid(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func TestAccSentinelAlertRuleMsSecurityIncident_withDisplayNameExcludeFilter(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_sentinel_alert_rule_ms_security_incident", "test")
	r := SentinelAlertRuleMsSecurityIncidentResource{}

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: r.displayNameExcludeFilter(data, "alert3"),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.displayNameExcludeFilter(data, "alert4"),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
		{
			Config: r.basic(data),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).ExistsInAzure(r),
			),
		},
		data.ImportStep(),
	})
}

func (t SentinelAlertRuleMsSecurityIncidentResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.AlertRuleID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.Sentinel.AlertRulesClient.Get(ctx, id.ResourceGroup, "Microsoft.OperationalInsights", id.WorkspaceName, id.Name)
	if err != nil {
		return nil, fmt.Errorf("reading Sentinel Alert Rule Ms Security Incident %q: %v", id, err)
	}

	rule, ok := resp.Value.(securityinsight.MicrosoftSecurityIncidentCreationAlertRule)
	if !ok {
		return nil, fmt.Errorf("reading Sentinel Alert Rule Ms Security Incident %q is not of MicrosoftSecurityIncidentCreationAlertRule kind", id)
	}

	return utils.Bool(rule.ID != nil), nil
}

func (r SentinelAlertRuleMsSecurityIncidentResource) basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_sentinel_alert_rule_ms_security_incident" "test" {
  name                       = "acctest-SentinelAlertRule-MSI-%d"
  log_analytics_workspace_id = azurerm_log_analytics_workspace.test.id
  product_filter             = "Microsoft Cloud App Security"
  display_name               = "some rule"
  severity_filter            = ["High"]
}
`, r.template(data), data.RandomInteger)
}

func (r SentinelAlertRuleMsSecurityIncidentResource) complete(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_sentinel_alert_rule_ms_security_incident" "test" {
  name                       = "acctest-SentinelAlertRule-MSI-%d"
  log_analytics_workspace_id = azurerm_log_analytics_workspace.test.id
  product_filter             = "Azure Security Center"
  display_name               = "updated rule"
  severity_filter            = ["High", "Low"]
  description                = "this is a alert rule"
  display_name_filter        = ["alert"]
}
`, r.template(data), data.RandomInteger)
}

func (r SentinelAlertRuleMsSecurityIncidentResource) requiresImport(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_sentinel_alert_rule_ms_security_incident" "import" {
  name                       = azurerm_sentinel_alert_rule_ms_security_incident.test.name
  log_analytics_workspace_id = azurerm_sentinel_alert_rule_ms_security_incident.test.log_analytics_workspace_id
  product_filter             = azurerm_sentinel_alert_rule_ms_security_incident.test.product_filter
  display_name               = azurerm_sentinel_alert_rule_ms_security_incident.test.display_name
  severity_filter            = azurerm_sentinel_alert_rule_ms_security_incident.test.severity_filter
}
`, r.basic(data))
}

func (r SentinelAlertRuleMsSecurityIncidentResource) alertRuleTemplateGuid(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

resource "azurerm_sentinel_alert_rule_ms_security_incident" "test" {
  name                       = "acctest-SentinelAlertRule-MSI-%d"
  log_analytics_workspace_id = azurerm_log_analytics_workspace.test.id
  product_filter             = "Microsoft Cloud App Security"
  display_name               = "some rule"
  severity_filter            = ["High"]
  alert_rule_template_guid   = "b3cfc7c0-092c-481c-a55b-34a3979758cb"
}
`, r.template(data), data.RandomInteger)
}

func (r SentinelAlertRuleMsSecurityIncidentResource) displayNameExcludeFilter(data acceptance.TestData, displayNameExcludeFilter string) string {
	return fmt.Sprintf(`
%s

resource "azurerm_sentinel_alert_rule_ms_security_incident" "test" {
  name                        = "acctest-SentinelAlertRule-MSI-%d"
  log_analytics_workspace_id  = azurerm_log_analytics_workspace.test.id
  product_filter              = "Microsoft Cloud App Security"
  display_name                = "some rule"
  severity_filter             = ["High"]
  display_name_filter         = ["alert1"]
  display_name_exclude_filter = ["%s"]
}
`, r.template(data), data.RandomInteger, displayNameExcludeFilter)
}

func (SentinelAlertRuleMsSecurityIncidentResource) template(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctestRG-sentinel-%d"
  location = "%s"
}

resource "azurerm_log_analytics_workspace" "test" {
  name                = "acctestLAW-%d"
  location            = azurerm_resource_group.test.location
  resource_group_name = azurerm_resource_group.test.name
  sku                 = "PerGB2018"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomInteger)
}
