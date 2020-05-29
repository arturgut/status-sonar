import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface APIResponse {
  site : {
    url: string,
    responseCode: number,
    durationInMs: number
  };
}

@Injectable({
  providedIn: 'root'
})
export class ApiResponseService {

  constructor( private http: HttpClient ) {}

  public getList() {
    return this.http.get<APIResponse>('http://localhost:8091/api/list');
  }
}
