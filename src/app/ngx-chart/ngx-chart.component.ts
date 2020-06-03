import { Component, NgModule, Input  } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { NgxChartsModule } from '@swimlane/ngx-charts';
import { ApiResponseService } from '../api-response.service';

@Component({
  selector: 'app-ngx-chart',
  templateUrl: './ngx-chart.component.html',
  styleUrls: ['./ngx-chart.component.css']
})
export class NgxChartComponent {

  public responseArrayProcess: any = [];
  public responseArrayTemp: any = [];
  public responseArrayNgxFormat: any = [];
  public tempArray: any = [];
  
  view: any[] = [1000, 300];

  // options
  legend: boolean = false;
  showLabels: boolean = true;
  animations: boolean = true;
  xAxis: boolean = true;
  yAxis: boolean = true;
  showYAxisLabel: boolean = true;
  showXAxisLabel: boolean = true;
  xAxisLabel: string = 'Duration in Ms';
  yAxisLabel: string = 'URL';
  timeline: boolean = true;
  rangeFillOpacity = 0.15
  colorScheme = 'fire';
  theme = 'dark';

  constructor(private urlfetchlist: ApiResponseService) {
    this.fetchUrls();
  }

  onClick(){
    this.fetchUrls();
  }

  fetchUrls(){
    this.urlfetchlist.getList().subscribe((response) => {
      this.responseArrayProcess = response;
      this.responseArrayNgxFormat = [];
      this.responseArrayProcess.forEach(element => {
        this.responseArrayNgxFormat.push({
          "name": element.url, "value": element.durationInMs 
        });
      });
      this.responseArrayNgxFormat = [...this.responseArrayNgxFormat];
    });
  }

  onSelect(data): void {}
  onActivate(data): void {}
  onDeactivate(data): void {}

}
