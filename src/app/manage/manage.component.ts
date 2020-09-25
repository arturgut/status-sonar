import { Component, OnInit, Input } from '@angular/core';
import { ApiResponseService } from '../api-response.service';

@Component({
  selector: 'app-manage',
  templateUrl: './manage.component.html',
  styleUrls: ['./manage.component.css']
})
export class ManageComponent implements OnInit {
  

  url: string = '';
  refreshRate: number;

  urlNew: string = 'https://your-site.com';
  refreshRateNew: number = 60;
  siteNameNew: string = 'Your site name';

  public accountObject: any;
  public urlList: any;
 
  constructor(private urlHandler: ApiResponseService) {
    this.fetchAccount();
  }

  onSubmit(){
    console.log("AccountObject: ", this.accountObject)
    this.updateAccount();
    this.reloadConfig();
  }

  onClick(){
    this.fetchUrls();
  }

  onRemove(key: string){
    delete this.accountObject.URLList[key]
  }

  onEdit() {

  }

  fetchAccount(){
    this.urlHandler.getAccount().subscribe((response) => {
      // Create array from object 
      this.accountObject = response;
    });
  }

  fetchUrls(){
    this.urlHandler.getList().subscribe((response) => {
      this.accountObject = response;
    });
  }

  onAdd(newSite: string, newUrl: string, newRefreshRate: number){
    console.log("Recived: ", newSite, newUrl, newRefreshRate)
    this.accountObject.URLList[newSite] =  {"url":newUrl,"responsecode":200, "durationinms":newRefreshRate,"timestamp":"2020-06-03T14:06:55.652Z" }
    console.log("Account object",this.accountObject)
  }

  updateAccount(){
    this.urlHandler.updateAccount(this.accountObject).subscribe((response) => {
      console.log("Reponse from API: ",response)
    }); 
  }

  reloadConfig() {
    this.urlHandler.reloadConfig().subscribe((response) => {
      console.log("Config reload requested.")
    });
  }

  ngOnInit(): void {
  }
}
