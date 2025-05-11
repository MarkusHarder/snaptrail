import { Component, OnDestroy, OnInit } from '@angular/core';

import { MenubarModule } from 'primeng/menubar';
import { MenuItem } from 'primeng/api';
import { Router, RouterModule } from '@angular/router';
import { CardModule } from 'primeng/card';
import { AuthService } from '../../services/auth.service';
import { Subscription } from 'rxjs';
import { ButtonModule } from 'primeng/button';

@Component({
  selector: 'app-menu-nav',
  imports: [MenubarModule, RouterModule, CardModule, ButtonModule],
  templateUrl: './menu-nav.component.html',
  styleUrl: './menu-nav.component.css',
})
export class MenuNavComponent implements OnInit, OnDestroy {
  items: MenuItem[] | undefined;
  loggedIn = false;
  authSub$: Subscription;

  constructor(
    private authService: AuthService,
    private router: Router,
  ) {
    this.authSub$ = this.authService.loggedIn$.subscribe((loggedIn) => {
      this.loggedIn = loggedIn;
      this.items = this.createItems();
    });
  }

  ngOnInit() {
    this.items = this.createItems();
  }

  createItems() {
    return [
      {
        label: 'Home',
        icon: 'pi pi-home',
        route: '/',
      },
      {
        label: 'Timeline',
        icon: 'pi pi-image',
        route: '/timeline',
      },
      {
        label: 'Sessions',
        icon: 'pi pi-folder-open',
        route: '/sessions',
        visible: !this.loggedIn,
      },
      {
        label: 'Sessions',
        icon: 'pi pi-folder-open',
        multi: this.loggedIn,
        visible: this.loggedIn,
        items: [
          {
            label: 'Manage',
            icon: 'pi pi-server',
            route: '/admin/sessions',
          },
          {
            label: 'Public',
            icon: 'pi pi-images',
            route: '/sessions',
          },
        ],
      },
    ];
  }

  logout() {
    this.authService.clearToken();
    this.router.navigate(['/admin/login']);
  }
  ngOnDestroy(): void {
    this.authSub$.unsubscribe();
  }
}
