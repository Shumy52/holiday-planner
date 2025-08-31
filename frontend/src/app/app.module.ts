import { APP_INITIALIZER, Injectable, NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AppComponent } from './app.component';
import { initKeycloak, keycloak } from './keycloak';


@Injectable()
export class TokenInterceptor {
  intercept(req:any,next:any){
    const token = (keycloak && keycloak.token) ? keycloak.token : null;
    return next.handle(token ? req.clone({ setHeaders:{ Authorization:`Bearer ${token}` }}) : req);
  }
}

@NgModule({
  imports: [BrowserModule, HttpClientModule, FormsModule, AppComponent],
  providers: [
    { provide: APP_INITIALIZER, useFactory: ()=>initKeycloak, multi:true },
    { provide: HTTP_INTERCEPTORS, useClass: TokenInterceptor, multi:true }
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
