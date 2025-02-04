package client

import (
	"github.com/auth0/go-auth0/management"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func flattenCustomSocialConfiguration(customSocial *management.ClientNativeSocialLogin) []interface{} {
	if customSocial == nil {
		return nil
	}

	m := map[string]interface{}{
		"apple": []interface{}{
			map[string]interface{}{
				"enabled": customSocial.GetApple().GetEnabled(),
			},
		},
		"facebook": []interface{}{
			map[string]interface{}{
				"enabled": customSocial.GetFacebook().GetEnabled(),
			},
		},
	}

	return []interface{}{m}
}

func flattenClientJwtConfiguration(jwt *management.ClientJWTConfiguration) []interface{} {
	if jwt == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"lifetime_in_seconds": jwt.GetLifetimeInSeconds(),
			"secret_encoded":      jwt.GetSecretEncoded(),
			"scopes":              jwt.GetScopes(),
			"alg":                 jwt.GetAlgorithm(),
		},
	}
}

func flattenClientRefreshTokenConfiguration(refreshToken *management.ClientRefreshToken) []interface{} {
	if refreshToken == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"rotation_type":                refreshToken.GetRotationType(),
			"expiration_type":              refreshToken.GetExpirationType(),
			"leeway":                       refreshToken.GetLeeway(),
			"token_lifetime":               refreshToken.GetTokenLifetime(),
			"infinite_token_lifetime":      refreshToken.GetInfiniteTokenLifetime(),
			"infinite_idle_token_lifetime": refreshToken.GetInfiniteIdleTokenLifetime(),
			"idle_token_lifetime":          refreshToken.GetIdleTokenLifetime(),
		},
	}
}

func flattenClientMobile(mobile *management.ClientMobile) []interface{} {
	if mobile == nil {
		return nil
	}

	m := map[string]interface{}{
		"android": nil,
		"ios":     nil,
	}

	if mobile.GetAndroid() != nil {
		m["android"] = []interface{}{
			map[string]interface{}{
				"app_package_name":         mobile.GetAndroid().GetAppPackageName(),
				"sha256_cert_fingerprints": mobile.GetAndroid().GetKeyHashes(),
			},
		}
	}

	if mobile.GetIOS() != nil {
		m["ios"] = []interface{}{
			map[string]interface{}{
				"team_id":               mobile.GetIOS().GetTeamID(),
				"app_bundle_identifier": mobile.GetIOS().GetAppID(),
			},
		}
	}

	return []interface{}{m}
}

func flattenClientAddons(addons *management.ClientAddons) []interface{} {
	if addons == nil {
		return nil
	}

	m := map[string]interface{}{
		"aws":                    flattenClientAddonAWS(addons.GetAWS()),
		"azure_blob":             flattenClientAddonAzureBlob(addons.GetAzureBlob()),
		"azure_sb":               flattenClientAddonAzureSB(addons.GetAzureSB()),
		"rms":                    flattenClientAddonRMS(addons.GetRMS()),
		"mscrm":                  flattenClientAddonMSCRM(addons.GetMSCRM()),
		"slack":                  flattenClientAddonSlack(addons.GetSlack()),
		"sentry":                 flattenClientAddonSentry(addons.GetSentry()),
		"echosign":               flattenClientAddonEchoSign(addons.GetEchoSign()),
		"egnyte":                 flattenClientAddonEgnyte(addons.GetEgnyte()),
		"firebase":               flattenClientAddonFirebase(addons.GetFirebase()),
		"newrelic":               flattenClientAddonNewRelic(addons.GetNewRelic()),
		"office365":              flattenClientAddonOffice365(addons.GetOffice365()),
		"salesforce":             flattenClientAddonSalesforce(addons.GetSalesforce()),
		"salesforce_api":         flattenClientAddonSalesforceAPI(addons.GetSalesforceAPI()),
		"salesforce_sandbox_api": flattenClientAddonSalesforceSandboxAPI(addons.GetSalesforceSandboxAPI()),
		"layer":                  flattenClientAddonLayer(addons.GetLayer()),
		"sap_api":                flattenClientAddonSAPAPI(addons.GetSAPAPI()),
		"sharepoint":             flattenClientAddonSharePoint(addons.GetSharePoint()),
		"springcm":               flattenClientAddonSpringCM(addons.GetSpringCM()),
		"wams":                   flattenClientAddonWAMS(addons.GetWAMS()),
		"zendesk":                flattenClientAddonZendesk(addons.GetZendesk()),
		"zoom":                   flattenClientAddonZoom(addons.GetZoom()),
		"sso_integration":        flattenClientAddonSSOIntegration(addons.GetSSOIntegration()),
		"samlp":                  flattenClientAddonSAML2(addons.GetSAML2()),
	}

	if addons.GetBox() != nil {
		m["box"] = []interface{}{map[string]interface{}{}}
	}

	if addons.GetCloudBees() != nil {
		m["cloudbees"] = []interface{}{map[string]interface{}{}}
	}

	if addons.GetConcur() != nil {
		m["concur"] = []interface{}{map[string]interface{}{}}
	}

	if addons.GetDropbox() != nil {
		m["dropbox"] = []interface{}{map[string]interface{}{}}
	}

	if addons.GetWSFED() != nil {
		m["wsfed"] = []interface{}{map[string]interface{}{}}
	}

	return []interface{}{m}
}

func flattenClientAddonAWS(addon *management.AWSClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"principal":           addon.GetPrincipal(),
			"role":                addon.GetRole(),
			"lifetime_in_seconds": addon.GetLifetimeInSeconds(),
		},
	}
}

func flattenClientAddonAzureBlob(addon *management.AzureBlobClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"account_name":       addon.GetAccountName(),
			"storage_access_key": addon.GetStorageAccessKey(),
			"container_name":     addon.GetContainerName(),
			"blob_name":          addon.GetBlobName(),
			"expiration":         addon.GetExpiration(),
			"signed_identifier":  addon.GetSignedIdentifier(),
			"blob_read":          addon.GetBlobRead(),
			"blob_write":         addon.GetBlobWrite(),
			"blob_delete":        addon.GetBlobDelete(),
			"container_read":     addon.GetContainerRead(),
			"container_write":    addon.GetContainerWrite(),
			"container_delete":   addon.GetContainerDelete(),
			"container_list":     addon.GetContainerList(),
		},
	}
}

func flattenClientAddonAzureSB(addon *management.AzureSBClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"namespace":    addon.GetNamespace(),
			"sas_key_name": addon.GetSASKeyName(),
			"sas_key":      addon.GetSASKey(),
			"entity_path":  addon.GetEntityPath(),
			"expiration":   addon.GetExpiration(),
		},
	}
}

func flattenClientAddonRMS(addon *management.RMSClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"url": addon.GetURL(),
		},
	}
}

func flattenClientAddonMSCRM(addon *management.MSCRMClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"url": addon.GetURL(),
		},
	}
}

func flattenClientAddonSlack(addon *management.SlackClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"team": addon.GetTeam(),
		},
	}
}

func flattenClientAddonSentry(addon *management.SentryClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"org_slug": addon.GetOrgSlug(),
			"base_url": addon.GetBaseURL(),
		},
	}
}

func flattenClientAddonEchoSign(addon *management.EchoSignClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"domain": addon.GetDomain(),
		},
	}
}

func flattenClientAddonEgnyte(addon *management.EgnyteClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"domain": addon.GetDomain(),
		},
	}
}

func flattenClientAddonFirebase(addon *management.FirebaseClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"secret":              addon.GetSecret(),
			"private_key_id":      addon.GetPrivateKeyID(),
			"private_key":         addon.GetPrivateKey(),
			"client_email":        addon.GetClientEmail(),
			"lifetime_in_seconds": addon.GetLifetimeInSeconds(),
		},
	}
}

func flattenClientAddonNewRelic(addon *management.NewRelicClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"account": addon.GetAccount(),
		},
	}
}

func flattenClientAddonOffice365(addon *management.Office365ClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"domain":     addon.GetDomain(),
			"connection": addon.GetConnection(),
		},
	}
}

func flattenClientAddonSalesforce(addon *management.SalesforceClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"entity_id": addon.GetEntityID(),
		},
	}
}

func flattenClientAddonSalesforceAPI(addon *management.SalesforceAPIClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"client_id":             addon.GetClientID(),
			"principal":             addon.GetPrincipal(),
			"community_name":        addon.GetCommunityName(),
			"community_url_section": addon.GetCommunityURLSection(),
		},
	}
}

func flattenClientAddonSalesforceSandboxAPI(addon *management.SalesforceSandboxAPIClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"client_id":             addon.GetClientID(),
			"principal":             addon.GetPrincipal(),
			"community_name":        addon.GetCommunityName(),
			"community_url_section": addon.GetCommunityURLSection(),
		},
	}
}

func flattenClientAddonLayer(addon *management.LayerClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"provider_id": addon.GetProviderID(),
			"key_id":      addon.GetKeyID(),
			"private_key": addon.GetPrivateKey(),
			"principal":   addon.GetPrincipal(),
			"expiration":  addon.GetExpiration(),
		},
	}
}

func flattenClientAddonSAPAPI(addon *management.SAPAPIClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"client_id":              addon.GetClientID(),
			"username_attribute":     addon.GetUsernameAttribute(),
			"token_endpoint_url":     addon.GetTokenEndpointURL(),
			"scope":                  addon.GetScope(),
			"service_password":       addon.GetServicePassword(),
			"name_identifier_format": addon.GetNameIdentifierFormat(),
		},
	}
}

func flattenClientAddonSharePoint(addon *management.SharePointClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"url":          addon.GetURL(),
			"external_url": addon.GetExternalURL(),
		},
	}
}

func flattenClientAddonSpringCM(addon *management.SpringCMClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"acs_url": addon.GetACSURL(),
		},
	}
}

func flattenClientAddonWAMS(addon *management.WAMSClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"master_key": addon.GetMasterkey(),
		},
	}
}

func flattenClientAddonZendesk(addon *management.ZendeskClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"account_name": addon.GetAccountName(),
		},
	}
}

func flattenClientAddonZoom(addon *management.ZoomClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"account": addon.GetAccount(),
		},
	}
}

func flattenClientAddonSSOIntegration(addon *management.SSOIntegrationClientAddon) []interface{} {
	if addon == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"name":    addon.GetName(),
			"version": addon.GetVersion(),
		},
	}
}

func flattenClientAddonSAML2(addon *management.SAML2ClientAddon) []interface{} {
	if addon == nil || addon.String() == "{}" {
		return nil
	}

	var logout interface{}

	if addon.GetLogout() != nil {
		logout = []interface{}{
			map[string]interface{}{
				"callback":    addon.GetLogout().GetCallback(),
				"slo_enabled": addon.GetLogout().GetSLOEnabled(),
			},
		}
	}

	return []interface{}{
		map[string]interface{}{
			"mappings":                           addon.GetMappings(),
			"audience":                           addon.GetAudience(),
			"recipient":                          addon.GetRecipient(),
			"create_upn_claim":                   addon.GetCreateUPNClaim(),
			"map_unknown_claims_as_is":           addon.GetMapUnknownClaimsAsIs(),
			"passthrough_claims_with_no_mapping": addon.GetPassthroughClaimsWithNoMapping(),
			"map_identities":                     addon.GetMapIdentities(),
			"signature_algorithm":                addon.GetSignatureAlgorithm(),
			"digest_algorithm":                   addon.GetDigestAlgorithm(),
			"issuer":                             addon.GetIssuer(),
			"destination":                        addon.GetDestination(),
			"lifetime_in_seconds":                addon.GetLifetimeInSeconds(),
			"sign_response":                      addon.GetSignResponse(),
			"name_identifier_format":             addon.GetNameIdentifierFormat(),
			"name_identifier_probes":             addon.GetNameIdentifierProbes(),
			"authn_context_class_ref":            addon.GetAuthnContextClassRef(),
			"typed_attributes":                   addon.GetTypedAttributes(),
			"include_attribute_name_format":      addon.GetIncludeAttributeNameFormat(),
			"binding":                            addon.GetBinding(),
			"signing_cert":                       addon.GetSigningCert(),
			"logout":                             logout,
		},
	}
}

func flattenClient(data *schema.ResourceData, client *management.Client) error {
	result := multierror.Append(
		data.Set("client_id", client.GetClientID()),
		data.Set("client_aliases", client.GetClientAliases()),
		data.Set("name", client.GetName()),
		data.Set("description", client.GetDescription()),
		data.Set("app_type", client.GetAppType()),
		data.Set("logo_uri", client.GetLogoURI()),
		data.Set("is_first_party", client.GetIsFirstParty()),
		data.Set("is_token_endpoint_ip_header_trusted", client.GetIsTokenEndpointIPHeaderTrusted()),
		data.Set("oidc_conformant", client.GetOIDCConformant()),
		data.Set("callbacks", client.GetCallbacks()),
		data.Set("allowed_logout_urls", client.GetAllowedLogoutURLs()),
		data.Set("allowed_origins", client.GetAllowedOrigins()),
		data.Set("allowed_clients", client.GetAllowedClients()),
		data.Set("grant_types", client.GetGrantTypes()),
		data.Set("organization_usage", client.GetOrganizationUsage()),
		data.Set("organization_require_behavior", client.GetOrganizationRequireBehavior()),
		data.Set("web_origins", client.GetWebOrigins()),
		data.Set("sso", client.GetSSO()),
		data.Set("sso_disabled", client.GetSSODisabled()),
		data.Set("cross_origin_auth", client.GetCrossOriginAuth()),
		data.Set("cross_origin_loc", client.GetCrossOriginLocation()),
		data.Set("custom_login_page_on", client.GetCustomLoginPageOn()),
		data.Set("custom_login_page", client.GetCustomLoginPage()),
		data.Set("form_template", client.GetFormTemplate()),
		data.Set("native_social_login", flattenCustomSocialConfiguration(client.GetNativeSocialLogin())),
		data.Set("jwt_configuration", flattenClientJwtConfiguration(client.GetJWTConfiguration())),
		data.Set("refresh_token", flattenClientRefreshTokenConfiguration(client.GetRefreshToken())),
		data.Set("encryption_key", client.GetEncryptionKey()),
		data.Set("addons", flattenClientAddons(client.GetAddons())),
		data.Set("mobile", flattenClientMobile(client.GetMobile())),
		data.Set("initiate_login_uri", client.GetInitiateLoginURI()),
		data.Set("signing_keys", client.SigningKeys),
		data.Set("client_metadata", client.GetClientMetadata()),
		data.Set("oidc_backchannel_logout_urls", client.GetOIDCBackchannelLogout().GetBackChannelLogoutURLs()),
		data.Set("require_pushed_authorization_requests", client.GetRequirePushedAuthorizationRequests()),
	)
	return result.ErrorOrNil()
}

func flattenClientForDataSource(data *schema.ResourceData, client *management.Client) error {
	result := multierror.Append(
		flattenClient(data, client),
		data.Set("client_secret", client.GetClientSecret()),
		data.Set("token_endpoint_auth_method", client.GetTokenEndpointAuthMethod()),
	)

	return result.ErrorOrNil()
}

func flattenClientGrant(data *schema.ResourceData, clientGrant *management.ClientGrant) error {
	result := multierror.Append(
		data.Set("client_id", clientGrant.GetClientID()),
		data.Set("audience", clientGrant.GetAudience()),
		data.Set("scopes", clientGrant.Scope),
	)

	return result.ErrorOrNil()
}
