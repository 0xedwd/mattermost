// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

//go:build enterprise

package enterprise

import (
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/account_migration"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/cluster"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/compliance"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/data_retention"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/elasticsearch"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/ldap"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/message_export"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/cloud"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/message_export/actiance_export"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/message_export/csv_export"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/message_export/global_relay_export"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/notification"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/oauth/google"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/oauth/office365"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/saml"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/oauth/openid"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/license"
	// Needed to ensure the init() method in the EE gets run
	_ "github.com/mattermost/enterprise/ip_filtering"
)
