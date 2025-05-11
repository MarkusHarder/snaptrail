import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MenuNavComponent } from './components/menu-nav/menu-nav.component';
import { CardModule } from 'primeng/card';
import { ToastModule } from 'primeng/toast';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, MenuNavComponent, CardModule, ToastModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent {
  title = 'snaptrail';
}
