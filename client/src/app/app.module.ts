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
import { AddFormComponent } from './components/Forms/add-form/add-form.component';
import { EditFormComponent } from './components/Forms/edit-form/edit-form.component';
import { SideNavComponent } from './layouts/side-nav/side-nav.component';
import { HeaderComponent } from './layouts/header/header.component';
import { TopThreeProductsComponent } from './components/Dashboard/top-three-products/top-three-products.component';
import { SalesByMonthComponent } from './components/Dashboard/sales-by-month/sales-by-month.component';
import { CardComponent } from './components/card/card.component';
import { HomeComponent } from './components/Home/home.component';
import { AddEditFormComponent } from './components/Forms/add-edit-form/add-edit-form.component';
import { SalesByCategoryComponent } from './components/Dashboard/sales-by-category/sales-by-category.component';
import { GamesComponent } from './components/ManageGame/games.component';

@NgModule({
  declarations: [
    AppComponent,
    GamesComponent,
    AddFormComponent,
    EditFormComponent,
    SideNavComponent,
    HeaderComponent,
    SalesByCategoryComponent,
    TopThreeProductsComponent,
    SalesByMonthComponent,
    CardComponent,
    HomeComponent,
    AddEditFormComponent
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
