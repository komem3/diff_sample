import { Injectable } from '@angular/core';
import { Message } from './model';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class DiffService {
  constructor(private http: HttpClient) {}

  CreateMessage(message: Message) {
    return this.http.post<Message>('/api/create', message);
  }
}
