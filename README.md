### Fork

Esta é uma versão em Português do Brasil do Documize.

Não há nenhuma funcionalidade adicional. O objetivo é possibilitar a avaliação
e uso da ferramenta por falantes da língua portuguesa até que o suporte a
múltiplos idiomas seja implementado no Documize.

Contribuições para a tradução são bem-vindas.

A tradução é feita diretamente no código-fonte e, por isso, os tradutores precisam ter noções de:
- Go
- Javascript
- Ember.js

This fork aims to provide a quick Brazilian portuguese translation of Documize while the official multilanguage support is not provided.


This effort is being done as part of an evaluation to choose the best suite to be used within SPU, a Brazilian public organization.

***


Documize is an open source modern, lightweight and comprehensive alternative to Confluence and other such solutions.

It's built with Golang + EmberJS and compiled down to a single executable binary for Linux, Windows and macOS.

All you need to provide is PostgreSQL, Microsoft SQL Server or any MySQL variant.

- Built for technical and non-technical users
- Designed to unify both customer-facing and internal documentation
- Organization through labels, spaces and categories
- No fee-charging marketplace

![Documize](screenshot-1.png "Documize")

## Latest Release

[Community Edition: v3.7.0](https://github.com/documize/community/releases)

[Enterprise Edition: v3.7.0](https://www.documize.com/downloads)

> *We provide frequent product updates for both cloud and self-hosted customers.*
>
> **Harvey Kandola, CEO/Founder @ Documize**

## OS Support

- Linux
- Windows
- macOS
- Raspberry Pi (using the ARM builds)

## Database Support

For all database types, Full-Text Search support (FTS) is mandatory.

- PostgreSQL (v9.6+)
- Microsoft SQL Server (2016+ with FTS)
- MySQL (v5.7.10+ and v8.0.12+)
- Percona (v5.7.16-10+)
- MariaDB (10.3.0+)

## Browser Support

- Chrome
- Firefox
- Safari
- Brave
- Vivaldi
- Opera
- Microsoft Edge (v42+)

## Technology Stack

- Go (v1.13.4)
- Ember JS (v3.12.0)

## Authentication Options

Besides email/password login, you can also authenticate via:

* LDAP
* Active Directory
* Red Hat Keycloak
* Central Authentication Service (CAS)

When using LDAP/Active Directory, you can enable dual-authentication with email/password.

## The Legal Bit

<https://documize.com>

This software (Documize Community Edition) is licensed under GNU AGPL v3 <http://www.gnu.org/licenses/agpl-3.0.en.html>. You can operate outside the AGPL restrictions by purchasing Documize Enterprise Edition and obtaining a commercial license by contacting <sales@documize.com>. Documize® is a registered trade mark of Documize Inc.

Documize uses other open source components and we acknowledge them in [NOTICES](NOTICES.md)
