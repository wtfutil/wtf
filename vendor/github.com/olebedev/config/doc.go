// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package config provides convenient access methods to configuration stored as
JSON or YAML.

Let's start with a simple YAML example:

    development:
      database:
        host: localhost
      users:
        - name: calvin
          password: yukon
        - name: hobbes
          password: tuna
    production:
      database:
        host: 192.168.1.1

We can parse it using ParseYaml(), which will return a *Config instance on
success:

    cfg, err := config.ParseYaml(yamlString)

An equivalent JSON configuration could be built using ParseJson():

    cfg, err := config.ParseJson(jsonString)

From now, we can retrieve configuration values using a path in dotted notation:

    // "localhost"
    host, err := cfg.String("development.database.host")

    // or...

    // "192.168.1.1"
    host, err := cfg.String("production.database.host")

Besides String(), other types can be fetched directly: Bool(), Float64(),
Int(), Map() and List(). All these methods will return an error if the path
doesn't exist, or the value doesn't match or can't be converted to the
requested type.

A nested configuration can be fetched using Get(). Here we get a new *Config
instance with a subset of the configuration:

    cfg, err := cfg.Get("development")

Then the inner values are fetched relatively to the subset:

    // "localhost"
    host, err := cfg.String("database.host")

For lists, the dotted path must use an index to refer to a specific value.
To retrieve the information from a user stored in the configuration above:

    // map[string]interface{}{ ... }
    user1, err := cfg.Map("development.users.0")
    // map[string]interface{}{ ... }
    user2, err := cfg.Map("development.users.1")

    // or...

    // "calvin"
    name1, err := cfg.String("development.users.0.name")
    // "hobbes"
    name2, err := cfg.String("development.users.1.name")

JSON or YAML strings can be created calling the appropriate Render*()
functions. Here's how we render a configuration like the one used in these
examples:

    cfg := map[string]interface{}{
        "development": map[string]interface{}{
            "database": map[string]interface{}{
                "host": "localhost",
            },
            "users": []interface{}{
                map[string]interface{}{
                    "name":     "calvin",
                    "password": "yukon",
                },
                map[string]interface{}{
                    "name":     "hobbes",
                    "password": "tuna",
                },
            },
        },
        "production": map[string]interface{}{
            "database": map[string]interface{}{
                "host": "192.168.1.1",
            },
        },
    }

    json, err := config.RenderJson(cfg)

    // or...

    yaml, err := config.RenderYaml(cfg)

This results in a configuration string to be stored in a file or database.

For more more convenience it can parse OS environment variables and command line arguments.

    cfg, err := config.ParseYaml(yamlString)
    cfg.Env()

    // or

    cfg.Flag()

We can also specify the order of parsing:

    cfg.Env().Flag()

    // or

    cfg.Flag().Env()

In case of OS environment all existing at the moment of parsing keys will be scanned in OS environment,
but in uppercase and the separator will be `_` instead of a `.`. If EnvPrefix() is used the given prefix
will be used to lookup the environment variable, e.g PREFIX_FOO_BAR will set foo.bar.
In case of flags separator will be `-`.
In case of command line arguments possible to use regular dot notation syntax for all keys.
For see existing keys we can run application with `-h`.

We can use unsafe method to get value:

  // ""
  cfg.UString("undefined.key")

  // or with default value
  unsafeValue := cfg.UString("undefined.key", "default value")

There is unsafe methods, like regular, but wuth prefix `U`.
*/
package config
