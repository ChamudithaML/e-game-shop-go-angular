// app/app.routes.ts
import { Routes } from '@angular/router';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { GamesComponent } from './pages/games/games.component';

export const routes: Routes = [
  {
    path: 'games',
    component:GamesComponent
  },
  {
    path:'dashboard',
    component:DashboardComponent
  }
];
