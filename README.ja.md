# spiffe-compliance-checker

[English](README.md) | **日本語**

SPIFFE 仕様の MUST / MUST NOT 条項を、与えられた artifact (SPIFFE ID 文字列、X.509-SVID 証明書、JWT-SVID トークン、Trust Bundle) に対して機械的にチェックする CLI。失敗時には仕様書名とセクション番号を出すので、そのまま spec 本文に飛んで根拠を確認できる。

SPIFFE プロジェクトは公式の conformance suite を提供していない。このツールは「artifact 単位で外から検証できる範囲」をカバーする。実装が workload attestation を正しく行っているか、鍵ローテーションを正しく回しているかなどの「動的な準拠性」は検証範囲外 (それは別問題)。

## なぜ

「SPIFFE 準拠」という言い回しは SPIRE / Istio / Cilium / 各社内製実装で頻出するが、何をもって準拠なのかは曖昧なまま使われがち。仕様は `spiffe/spiffe` の 8 本の markdown に散らばっていて、MUST 句は各所に埋め込まれている。1 度読めば理解できる量だが、毎回証明書を渡されるたびに手で確認するのは現実的ではない。

`scc` は仕様の MUST / MUST NOT を直接読み下した assertion を、artifact に対して一発で適用する。

## 使い方

```text
scc id        <spiffe-id-string>
scc x509-svid <cert.pem | cert.der>
scc jwt-svid  <token>
scc bundle    <bundle.json>
```

各サブコマンドは assertion 1 件につき 1 行を出力する。1 件でも MUST 句が落ちれば exit code は非ゼロ。

```text
$ scc id 'spiffe://Example.com/web-fe'
FAIL  SPIFFE-ID.md §2.1   trust domain MUST be lowercase
PASS  SPIFFE-ID.md §2     scheme is "spiffe"
PASS  SPIFFE-ID.md §2.1   trust domain non-empty
...
```

## カバレッジ

| 仕様                                  | `scc` がチェックするもの                                                                            |
| ------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `SPIFFE-ID.md`                        | scheme、trust domain の charset / 長さ / case、path segment、URI 全長、query/fragment 不在          |
| `X509-SVID.md`                        | URI SAN 個数、leaf/CA の Basic Constraints、Key Usage critical + 各 flag、EKU、leaf SPIFFE ID 規約  |
| `JWT-SVID.md`                         | `alg` whitelist、JWS Compact Serialization、`sub`/`aud`/`exp` の存在、`sub` の SPIFFE ID 妥当性     |
| `SPIFFE_Trust_Domain_and_Bundle.md`   | JWKS shape、key ごとの `kty`/`use`、`spiffe_sequence` / `spiffe_refresh_hint` の型、x509 の `x5c`   |

カバーしないもの: 実際に動いている `Workload API` endpoint の振る舞い (別リポで扱う想定。Agent が必要)、Federation endpoint の信頼関係、特定 Trust Bundle に対する署名検証。1 つ目はロードマップ、残り 2 つは意図的にスコープ外。

## インストール

```bash
go install github.com/0-draft/spiffe-compliance-checker/cmd/scc@latest
```

Go 1.22 以降が必要。バイナリは追加ランタイム依存なし。

## ライセンス

Apache-2.0。詳細は `LICENSE`。
