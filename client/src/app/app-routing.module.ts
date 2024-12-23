import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { GamesComponent } from './pages/games/games.component';
import { PracticePageComponent } from './pages/practice-page/practice-page.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { HomeComponent } from './pages/home/home.component';
import { DynamicFormComponent } from './pages/dynamic-form/dynamic-form.component';
import { DynamicFormTwoComponent } from './pages/dynamic-form-two/dynamic-form-two.component';

const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'dash', component: DashboardComponent },
  { path: 'games', component: GamesComponent },
  { path: 'prac', component: PracticePageComponent },
  { path: 'dyn', component: DynamicFormComponent },
  { path: 'dyn2', component: DynamicFormTwoComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
