import { Component, OnInit, Input } from '@angular/core';
import { ApiResponseService } from '../api-response.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {


     public responseArray: any;
//   public responseObject = [
//     {
//        "url": "http://bbc.com",
//        "responseCode": 200,
//        "durationInMs": 232
//     },
//     {
//        "url": "http://dell.com",
//        "responseCode": 200,
//        "durationInMs": 434
//     },
//     {
//        "url": "http://google.com",
//        "responseCode": 200,
//        "durationInMs": 232
//     }
//  ];
  public url: string;
  public responseCode: number;
  public durationInMs: number;
  
  constructor(private urlfetchlist: ApiResponseService) {
    this.fetchUrls()
  }

  onClick(){
    this.fetchUrls()
  }

  fetchUrls(){
    this.urlfetchlist.getList().subscribe((response) => {
      console.log(response);
      this.responseArray = response;
      // console.log(this.responseObject);
      // this.responseCode = response.site.responseCode; 
      // this.durationInMs = response.site.durationInMs;
      // this.responseCode = response.site.responseCode;
      // console.log(response.site.url);
    });
  }

  ngOnInit(): void {
  }

}
