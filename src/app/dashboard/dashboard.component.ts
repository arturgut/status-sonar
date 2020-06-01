import { Component, OnInit, Input } from '@angular/core';
import { ApiResponseService } from '../api-response.service';
import { NgxChartComponent } from '../ngx-chart/ngx-chart.component';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {

  public responseArray: any;
 
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
    });
  }

  ngOnInit(): void {
  }

}
