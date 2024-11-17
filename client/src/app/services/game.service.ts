import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';

import { GameResponse } from '../models/game-response.model'; 
import { Game } from '../models/game.model';

@Injectable({
  providedIn: 'root'
})
export class GameService {

  constructor(private http: HttpClient) { }

  getAll(): Observable<GameResponse> {
    return this.http.get<GameResponse>(environment.proxyUrl);
  }

  create(data: Game): Observable<any> {
    return this.http.post(environment.proxyUrlAdd, data);
  }

  update(id: any, data: any): Observable<any> {
    return this.http.put(`${environment.baseUrl}/${id}`, data);
  }

  delete(id: any): Observable<any> {
    return this.http.delete(`${environment.baseUrl}/${id}`);
  }

}