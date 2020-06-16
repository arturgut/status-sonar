import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface URL {
  site : {
    url: string,
    responseCode: number,
    durationInMs: number
  };
}

interface Account {
  accountname: string, 
  urllist: URL[]
}

@Injectable({
  providedIn: 'root'
})
export class ApiResponseService {

  account: Account


  constructor( private http: HttpClient ) {}

  public getAccount() {
    return this.http.get<Account>('http://localhost:8092/api/account/get',{
      params: {
        account: 'default'
      }
    });
  }

  public getList() {
    return this.http.get<URL>('http://localhost:8091/api/list',{
      params: {
        account: 'default'
      }
    });
  }

  public updateAccount(account: Account ) {

    return this.http.post<URL>('http://localhost:8092/api/account/update', account )

  }

}
