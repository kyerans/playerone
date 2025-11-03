# playerone
implement key system use for encryption media

1. Encoder/Packager:
   Generate ContentKey Kc + KID
   Encrypt video with Kc
   Embed KID in manifest

2. DRM Server:
   Store KID → Kc mapping

3. Player:
   Read manifest → get KID
   POST LicenseRequest { kids: [KID] }

4. DRM Server:
   Lookup KID → return LicenseResponse { keys: [{k,kid,kty}] }

5. Player:
   Use key to decrypt video