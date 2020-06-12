import { Component, OnInit } from '@angular/core';
import { ApiResponseService } from '../api-response.service';
import { element } from 'protractor';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  public responseArray: any;
  public urlList: any;
  // public responseArrayProcess: any;
  // public responseArrayNgxFormat = [];

  constructor(private urlfetchlist: ApiResponseService) {
    this.fetchUrls();
  }

  onClick(){
    this.fetchUrls();
  }

  fetchUrls(){
    this.urlfetchlist.getList().subscribe((response) => {
      this.responseArray = response;
      
    });
  }

  ngOnInit(): void {
  }

}
