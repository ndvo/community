// Copyright 2016 Documize Inc. <legal@documize.com>. All rights reserved.
//
// This software (Documize Community Edition) is licensed under
// GNU AGPL v3 http://www.gnu.org/licenses/agpl-3.0.en.html
//
// You can operate outside the AGPL restrictions by purchasing
// Documize Enterprise Edition and obtaining a commercial license
// by contacting <sales@documize.com>.
//
// https://documize.com

import { inject as service } from '@ember/service';
import RSVP from 'rsvp';
import AuthenticatedRouteMixin from 'ember-simple-auth/mixins/authenticated-route-mixin';
import Route from '@ember/routing/route';

export default Route.extend(AuthenticatedRouteMixin, {
	orgService: service('organization'),
	appMeta: service(),
	session: service(),

	beforeModel() {
		if (!this.get("session.isAdmin")) {
			this.transitionTo('auth.login');
		}
	},

	model() {
		let orgId = this.get("appMeta.orgId");

		if (this.get("session.isGlobalAdmin")) {
			return RSVP.hash({
				jira: this.get('orgService').getOrgSetting(orgId, 'jira'),
				trello: this.get('orgService').getGlobalSetting('SECTION-TRELLO')
			});
		} else {
			return RSVP.hash({
				jira: this.get('orgService').getOrgSetting(orgId, 'jira'),
				trello: { appKey: '' }
			});
		}
	},

	activate() {
		this.get('browser').setTitle('Integrações');
	}
});
