import { Injectable } from '@angular/core';
import { MessageReq, MessageResp } from './model';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class DiffService {
  constructor(private http: HttpClient) {}

  CreateMessage(message: MessageReq) {
    return this.http.post<MessageResp>('/api/message', message);
  }

  UpdateMessage(id: string, message: MessageReq) {
    return this.http.patch<MessageResp>(`/api/message/${id}`, message);
  }
}
