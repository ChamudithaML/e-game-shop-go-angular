import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

import { environment } from '../../environments/environment';
import { SaleResponse } from '../models/sale-response.model';

@Injectable({
  providedIn: 'root'
})
export class SalesService {

  constructor(private http: HttpClient) { }

  getAll(): Observable<SaleResponse> {
    return this.http.get<SaleResponse>(environment.saleUrl);
  }
}
