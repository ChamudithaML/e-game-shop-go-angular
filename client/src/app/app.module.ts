import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppRoutingModule } from './app-routing.module';


import { TableModule } from 'primeng/table';
import { DialogModule } from 'primeng/dialog'; 
import { DynamicDialogModule } from 'primeng/dynamicdialog';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card'; 
import { DialogService } from 'primeng/dynamicdialog';

import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { ChartModule } from 'angular-highcharts';

import { AppComponent } from './app.component';
import { GamesComponent } from './pages/games/games.component';
import { AddFormComponent } from './components/Forms/add-form/add-form.component';
import { EditFormComponent } from './components/Forms/edit-form/edit-form.component';
import { SideNavComponent } from './layouts/side-nav/side-nav.component';
import { HeaderComponent } from './layouts/header/header.component';
import { MainComponent } from './pages/main/main.component';
import { SalesByCategoryComponent } from './dashboard/sales-by-category/sales-by-category.component';
import { TopThreeProductsComponent } from './dashboard/top-three-products/top-three-products.component';
import { SalesByMonthComponent } from './dashboard/sales-by-month/sales-by-month.component';
import { CardComponent } from './components/card/card.component';
import { PracticePageComponent } from './pages/practice-page/practice-page.component';
import { TempDialogComponent } from './components/practice-components/temp-dialog/temp-dialog.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { HomeComponent } from './pages/home/home.component';
import { AddEditFormComponent } from './components/Forms/add-edit-form/add-edit-form.component';
import { DynamicFormComponent } from './pages/dynamic-form/dynamic-form.component';
import { DynamicFormTwoComponent } from './pages/dynamic-form-two/dynamic-form-two.component';

@NgModule({
  declarations: [
    AppComponent,
    GamesComponent,
    AddFormComponent,
    EditFormComponent,
    SideNavComponent,
    HeaderComponent,
    MainComponent,
    SalesByCategoryComponent,
    TopThreeProductsComponent,
    SalesByMonthComponent,
    CardComponent,
    TempDialogComponent,
    PracticePageComponent,
    DashboardComponent,
    HomeComponent,
    AddEditFormComponent,
    DynamicFormComponent,
    DynamicFormTwoComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    TableModule,
    DialogModule,
    BrowserModule,
    ButtonModule,
    BrowserAnimationsModule,
    FontAwesomeModule,
    ChartModule,
    CardModule,
    DynamicDialogModule,
    ReactiveFormsModule
  ],
  providers: [DialogService],
  bootstrap: [AppComponent]
})
export class AppModule { }
