# Version 0.1 - 'Simple Dashboard'   
- [x] UI - Dashboard Mockup 
- [x] UI - Generate Menu and Dashboard Components 
- [x] API - Create Status Sonar API microservice
- [x] API - Refactor mapToJSON() to output Array rather than Object
- [x] UI - Add routing 
- [x] UI - Add Service HTTP Client connection to API  
- [x] UI - Add *ngFor to list URL's from API 
- [x] UI - Add ngx-charts suppport with real time display fetched from api 
- [x] API - Add timestamp to /list 

# Version 0.2 - 'URL Manage'
- [x] API - Move urlchecker source code to sonar-status repo 
- [x] API - Add MongoDB driver
- [x] API - Define Account BSON structure
- [x] API - Add GET api/account/list endpoint
- [x] API - Add GET api/account/add endpoint
- [x] API - Add GET api/account/update endpoint
- [x] UI - Update dashboard to display data from account endpoint 
- [x] UI - Add 'URL Manage' component
- [x] API - Update urlcheck service to load list of urls from MongoDB
- [x] API - URLService - Add reload configuration endpoint
- [x] API - URLService - Reload configuration every 30s
- [ ] UI - Clean up 'URL Manage' view 
- [ ] UI + API - Hookup UI MANAGE to API

# Version 0.3 - 'User Accounts Administration'
- [ ] UI - Add account admin page 
- [ ] API - Add api/user/add, 
- [ ] API - Add api/user/remove, 
- [ ] API - Add api/user/list, 
- [ ] API - Add api/user/update

# Version 0.4 'User Authentication' 
- [ ] API - User authentication
- [ ] UI - User authentication

# Version 0.5 'CICD' 
- [ ] Create Docker 
- [ ] Create Jenkinsfile (Build + Test + Push to DockerHub Registry)
- [ ] Create Kubernetes Deployment   

# Version 0.6 'Enhanced Sonar' 
- [ ] Update API services to export internal metrics in prometheus format
- [ ] API - Add metrics database (TSDB or Prometheus) 
- [ ] API & UI - SSL validation (show for how long the certificate will be valid for) 
- [ ] UI - Status color changes depends on status (red for HTTP 5xx, Orange for HTTP 4xx, Green for HTTP 2xx)

# Version 0.6 'Multiregion Status Sonar' 
- [ ] API - Add support for multiregion checking 
