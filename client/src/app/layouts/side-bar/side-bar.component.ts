import { Component, OnInit } from '@angular/core';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { RouterModule } from '@angular/router';

import {
    faDashboard,
    faLocation,
    faShop,
    faBox,
    faMoneyBill,
    faChartBar,
    faContactBook,
    faHand,
} from '@fortawesome/free-solid-svg-icons';

@Component({
    selector: 'app-side-bar',
    standalone: true,
    imports: [FontAwesomeModule, RouterModule],
    templateUrl: './side-bar.component.html',
    styleUrl: './side-bar.component.scss'
})

export class SideBarComponent implements OnInit {

    faDashboard = faDashboard;
    faLocation = faLocation;
    faShop = faShop;
    faBox = faBox;
    faMoneyBill = faMoneyBill;
    faChartBar = faChartBar;
    faContactBook = faContactBook;
    faHand = faHand;

    constructor() { }

    ngOnInit(): void {
    }
}