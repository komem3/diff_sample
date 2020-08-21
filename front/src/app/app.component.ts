import { Component } from '@angular/core';
import { DiffService } from './diff.service';
import { Observable } from 'rxjs';
import { Message } from './model';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  input: Message = { message: '', id: '', diff: '' };
  diff: string;

  constructor(private diffService: DiffService) {}

  submit() {
    if (this.input.id) {
      this.diffService.UpdateMessage(this.input).subscribe((r) => {
        this.input.id = r.id;
        this.diff = decodeURI(r.diff);
      });
    } else {
      this.diffService.CreateMessage(this.input).subscribe((r) => {
        this.input.id = r.id;
        this.diff = r.message;
      });
    }
  }
}
