# spiffe-compliance-checker

**English** | [日本語](README.ja.md)

Static compliance checker for SPIFFE artifacts. Pass it a SPIFFE ID, an X.509-SVID certificate, a JWT-SVID token, or a SPIFFE trust bundle, and it tells you which MUST / MUST NOT clauses from the SPIFFE spec the artifact satisfies or violates. Each failure cites the spec document and section so you can jump straight to the source of truth.

The SPIFFE project does not ship an official conformance suite. This tool fills part of that gap for the parts that are checkable from outside: the artifacts themselves. It does not check whether an implementation correctly attests workloads or rotates keys — that requires runtime observation.

## Why

"This is SPIFFE-compliant" gets repeated across SPIRE, Istio, Cilium, and a growing list of in-house implementations, but the claim is rarely backed by a concrete checklist. The spec lives across eight markdown files in `spiffe/spiffe`, with MUST clauses scattered throughout. Reading them once is fine. Re-deriving the checklist every time someone hands you a cert is not.

`scc` reads the MUST/MUST NOT requirements directly from the spec and applies them as one-shot assertions on the artifact you give it.

## Usage

```text
scc id        <spiffe-id-string>
scc x509-svid <cert.pem | cert.der>
scc jwt-svid  <token>
scc bundle    <bundle.json>
```

Each subcommand prints one line per assertion. Exit code is non-zero if any MUST clause fails.

```text
$ scc id 'spiffe://Example.com/web-fe'
FAIL  SPIFFE-ID.md §2.1   trust domain MUST be lowercase
PASS  SPIFFE-ID.md §2     scheme is "spiffe"
PASS  SPIFFE-ID.md §2.1   trust domain non-empty
...
```

## Scope

| Spec                                  | What `scc` checks                                                                                |
| ------------------------------------- | ------------------------------------------------------------------------------------------------ |
| `SPIFFE-ID.md`                        | scheme, trust domain charset/length/case, path segments, total length, query/fragment exclusion  |
| `X509-SVID.md`                        | URI SAN count, leaf/CA Basic Constraints, Key Usage critical+flags, EKU, leaf SPIFFE ID rules    |
| `JWT-SVID.md`                         | `alg` whitelist, JWS Compact Serialization, `sub`/`aud`/`exp` presence, SPIFFE ID in `sub`       |
| `SPIFFE_Trust_Domain_and_Bundle.md`   | JWKS shape, `kty`/`use` per key, `spiffe_sequence`/`spiffe_refresh_hint` types, `x5c` for x509   |

What it does not check: live `Workload API` endpoint behavior (different repo concern, needs a running Agent), Federation endpoint trust, signature validity against a specific Trust Bundle. The first is on the roadmap; the latter two are deliberately out of scope.

## Install

```bash
go install github.com/0-draft/spiffe-compliance-checker/cmd/scc@latest
```

Requires Go 1.22 or newer. The binary has no runtime dependencies.

## License

Apache-2.0. See `LICENSE`.
