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

package database

import (
	"fmt"
	"sort"
	"strings"
	"github.com/documize/community/core/env"
	"github.com/documize/community/server/web"
)

// Scripts holds all .SQL files for all supported database providers.
type Scripts struct {
	MySQL      []Script
	PostgreSQL []Script
	SQLServer  []Script
	CustomMySQL      []Script
	CustomPostgreSQL []Script
	CustomSQLServer  []Script
}

// Script holds SQL script and it's associated version number.
type Script struct {
	Version int
	Script  []byte
}

// LoadScripts returns .SQL scripts for supported database providers.
func LoadScripts() (s Scripts, err error) {
	assetDir := "bindata/scripts"

	// MySQL
	s.MySQL, err = loadFiles(fmt.Sprintf("%s/mysql", assetDir))
	if err != nil {
		return
	}
	// PostgreSQL
	s.PostgreSQL, err = loadFiles(fmt.Sprintf("%s/postgresql", assetDir))
	if err != nil {
		return
	}
	// PostgreSQL
	s.SQLServer, err = loadFiles(fmt.Sprintf("%s/sqlserver", assetDir))
	if err != nil {
		return
	}

	return s, nil
}

// LoadCustom Scripts returns custom .SQL scripts for supported database providers.
func LoadCustomScripts(currentVersion int) (s Scripts, err error) {
	assetDir := "bindata/scripts"

	// No problem is files are not found, even if directories do not exist.
	// MySQL
	s.CustomMySQL, _ = loadFiles(fmt.Sprintf("%s/mysql/%d", assetDir, currentVersion))
	// PostgreSQL
	s.CustomPostgreSQL, _ = loadFiles(fmt.Sprintf("%s/postgresql/%d", assetDir, currentVersion))
	// PostgreSQL
	s.CustomSQLServer, _ = loadFiles(fmt.Sprintf("%s/sqlserver/%d", assetDir, currentVersion))

	return s, nil
}

// SpecificScripts returns SQL scripts for current databasse provider.
func SpecificScripts(runtime *env.Runtime, all Scripts) (s []Script) {
	switch runtime.StoreProvider.Type() {
	case env.StoreTypeMySQL, env.StoreTypeMariaDB, env.StoreTypePercona:
		return all.MySQL
	case env.StoreTypePostgreSQL:
		return all.PostgreSQL
	case env.StoreTypeSQLServer:
		return all.SQLServer
	}

	return
}

// CustomSpecificScripts returns custom SQL scripts for current databasse provider.
func CustomSpecificScripts(runtime *env.Runtime, all Scripts) (s []Script) {
	switch runtime.StoreProvider.Type() {
	case env.StoreTypeMySQL, env.StoreTypeMariaDB, env.StoreTypePercona:
		return all.CustomMySQL
	case env.StoreTypePostgreSQL:
		return all.CustomPostgreSQL
	case env.StoreTypeSQLServer:
		return all.CustomSQLServer
	}

	return
}

// loadFiles returns all SQL scripts in specified folder as [][]byte.
func loadFiles(path string) (b []Script, err error) {
	buf := []byte{}
	scripts, err := web.AssetDir(path)
	if err != nil {
		return
	}
	sort.Strings(scripts)
	for _, file := range scripts {
		if strings.HasSuffix(file, ".sql") {
			buf, err = web.Asset(fmt.Sprintf("%s/%s", path, file))
			if err != nil {
				return
			}

			b = append(b, Script{Version: extractVersionNumber(file), Script: buf})
		} 
	}

	return b, nil
}
