# sanitizer

Quick sanitizer that replaces PII, currently IPv4, IPv6 and host/domain names.

#### Show possible PII

```
./sanitizer show test-parsing-issue-5150.log
```

#### Generate PII replacement JSON configuration
Convenient when you need to preview and edit replacements first

```
./sanitizer gen test-parsing-issue-5150.log > replacement.json
```

#### Sanitize without preview/config
If you are feeling brave you could just get the sanitized content without config and hope for the best

```
./sanitizer san test-parsing-issue-5150.log > test-parsing-issue-5150.log.sanitized
```

#### Sanitize with a replacement configuration file
This option is convenient when you generate the replacement file first, review the replacements (might need to modify further or remove some of the assumed replacements) and then apply these replacements for sanitization.
```
./sanitizer san test-parsing-issue-5150.log replacement.json > test-parsing-issue-5150.log.sanitized
```

