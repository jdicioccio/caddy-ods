package ods

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsods "github.com/jdicioccio/libdns-ods"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *libdnsods.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.ods",
		New: func() caddy.Module { return &Provider{new(libdnsods.Provider)} },
	}
}

// TODO: This is just an example. Useful to allow env variable placeholders; update accordingly.
// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.Host = caddy.NewReplacer().ReplaceAll(p.Provider.Host, "")
	p.Provider.User = caddy.NewReplacer().ReplaceAll(p.Provider.User, "")
	p.Provider.Pass = caddy.NewReplacer().ReplaceAll(p.Provider.Pass, "")

	return nil
}

// TODO: This is just an example. Update accordingly.
// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// providername [<api_token>] {
//     api_token <api_token>
// }
//
// **THIS IS JUST AN EXAMPLE AND NEEDS TO BE CUSTOMIZED.**
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.Host = d.Val()
		}
		if d.NextArg() {
			p.Provider.User = d.Val()
		}
		if d.NextArg() {
			p.Provider.Pass = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "host":
				if p.Provider.Host != "" {
					return d.Err("Host already set")
				}
				if d.NextArg() {
					p.Provider.Host = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "user":
				if p.Provider.User != "" {
					return d.Err("User already set")
				}
				if d.NextArg() {
					p.Provider.User = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "pass":
				if p.Provider.Pass != "" {
					return d.Err("Pass already set")
				}
				if d.NextArg() {
					p.Provider.Pass = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.Host == "" {
		p.Provider.Host = "ratbox.int.ods.org"
	}
	if p.Provider.User == "" {
		return d.Err("missing User")
	}
	if p.Provider.Pass == "" {
		return d.Err("missing Pass")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
