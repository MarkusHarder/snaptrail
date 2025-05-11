import { Routes } from '@angular/router';
import { LandingPageComponent } from './components/landing-page/landing-page.component';
import { AdminLoginComponent } from './components/admin-login/admin-login.component';
import { TimelineComponent } from './components/timeline/timeline.component';
import { SessionsComponent } from './components/sessions/sessions.component';
import { AdminLayoutComponent } from './components/admin-layout/admin-layout.component';
import { authGuard } from './guards/auth.guard';
import { PublicSessionsComponent } from './components/public-sessions/public-sessions.component';
import { AdminUserPageComponent } from './components/admin-user-page/admin-user-page.component';

export const routes: Routes = [
  { path: '', component: LandingPageComponent },
  {
    path: 'admin',
    component: AdminLayoutComponent,
    children: [
      {
        path: 'login',
        component: AdminLoginComponent,
      },
      {
        path: 'sessions',
        component: SessionsComponent,
        canActivate: [authGuard],
      },
      {
        path: 'user',
        component: AdminUserPageComponent,
        canActivate: [authGuard],
      },
    ],
  },
  { path: 'timeline', component: TimelineComponent },
  { path: 'sessions', component: PublicSessionsComponent },
];
