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

import $ from 'jquery';
import Route from '@ember/routing/route';

export default Route.extend({
	beforeModel() {
		let pwd = document.head.querySelector("[property=dbhash]").content;
		if (pwd.length === 0 || pwd === "{{.DBhash}}") {
			// don't allow access to this page if we are not in setup mode
			this.transitionTo('auth.login');
		}
	},

	model() {
		let pwd = document.head.querySelector("[property=dbhash]").content;

		return {
			dbname: document.head.querySelector("[property=dbname]").content,
			dbhash: pwd,
			title: "",
			message: "Esta instância do documize conctém toda a documentação da nossa equipe",
			allowAnonymousAccess: false,
			firstname: "",
			lastname: "",
			email: "",
			password: pwd,
			activationKey: '',
			edition: document.head.querySelector("[property=edition]").content
		};
	},

	activate() {
		$('body').addClass('background-color-theme-100');
		document.title = "Bem-vindo ao instalador do Documize";
	},

	deactivate() {
		$('body').removeClass('background-color-theme-100');
	}
});
