import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FullCalendarModule } from '@fullcalendar/angular';
import { CalendarOptions } from '@fullcalendar/core';
import dayGridPlugin from '@fullcalendar/daygrid';
import interactionPlugin from '@fullcalendar/interaction';
import { HttpClient } from '@angular/common/http';

@Component({
  selector:'app-calendar',
  standalone: true,
  imports: [CommonModule, FullCalendarModule],
  template:`<full-calendar [options]="options"></full-calendar>`
})
export class CalendarComponent implements OnInit {
  options: CalendarOptions = {
    plugins: [dayGridPlugin, interactionPlugin],
    initialView: 'dayGridMonth',
    events: []
  };
  constructor(private http:HttpClient){}
  ngOnInit(){
    this.http.get<any[]>('/api/v1/vacations/mine').subscribe(vacs=>{
      this.options = {
        ...this.options,
        events: vacs.map(v=>({ title:`${v.status} (${v.totalDays}z)`,
          start:v.startDate, end:v.endDate }))
      };
    });
  }
}
