1. Swagger
2. Chain middleware
3. Check tenant permissions
   - Always do this at core layer to prevent logic duplication in each adapter.
4. App config
5. Connection context
6. Idiomatic package names
7. Core "convenience" methods to get things by string ID
   - Much less code in controllers using path ID params
8. OrgGroupFeatureFlag
   - Seed in CLI
   - Move upsert to Core
   - `PUT /org/{orgSlug}/app/{appSlug}/flag/{flagID}/group-flag` to replace
