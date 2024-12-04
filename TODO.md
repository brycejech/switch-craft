1. Swagger
2. Chain middleware
3. Check tenant permissions
   - Always do this at core layer to prevent logic duplication in each adapter.
4. Eliminate concept of "global account"
   - Will simplify logic greatly if all accounts must be tied to an org
   - Global org management can come much much later
   - Org create should be multiple steps
     - No auth
     - Create account with no org
     - Create org with owner newAccount.ID
     - Update account.OrgID to be newOrg.ID
