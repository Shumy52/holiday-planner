import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
@Component({
  selector:'app-request-form',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl:'./request-form.component.html'
})
export class RequestFormComponent {
  start=''; end=''; msg='';
  constructor(private http:HttpClient){}
  submit(){
    this.http.post('/api/v1/vacations',{start:this.start,end:this.end})
      .subscribe({
        next: _ => this.msg='Trimis!',
        error: e => this.msg = 'Eroare: '+(e.error?.error||e.message)
      });
  }
}
