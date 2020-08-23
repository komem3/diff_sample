import { Component } from '@angular/core';
import { DiffService } from './diff.service';
import { Observable } from 'rxjs';
import { MessageReq } from './model';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  input: string = '';
  id: string = '';
  diff: string = '';

  constructor(private diffService: DiffService) {}

  onClickCheck() {
    this.diffService.CreateMessage({ message: this.input }).subscribe((r) => {
      this.diff = r.message;
      this.id = r.id;
    });
  }
}
