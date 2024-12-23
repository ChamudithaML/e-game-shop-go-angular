
import { Component, OnInit } from '@angular/core';
import { Chart } from 'angular-highcharts';
import { Sales } from '../../../models/sales.model';
import { SalesService } from '../../../services/sales.service';

@Component({
  selector: 'app-sales-by-category',
  templateUrl: './sales-by-category.component.html',
  styleUrls: ['./sales-by-category.component.scss']
})
export class SalesByCategoryComponent implements OnInit {

  titles: string[] = []
  stocks: number[] = []
  sales: number[] = []
  saleSeries: Sales[] = []
  chart: Chart | undefined;

  constructor(private saleService: SalesService) { }

  retrieveSales(): void {
    this.saleService.getAll().subscribe(response => {

      // when the incoming data can be undefined that has to be handled by adding ! or ?? []
      this.titles = response.titles!;
      this.stocks = response.stocks ?? [];
      this.sales = response.sales!;
      this.generateChartSeries(this.titles, this.sales);

      this.initializeChart();
      // console.log(this.sales)
    }, error => {
      console.error("Error retrieving games:", error);
    });
  }

  getRandomColor(): string {
    const letters = '0123456789ABCDEF';
    let color = '#';
    for (let i = 0; i < 6; i++) {
      color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
  }

  generateChartSeries(gameNames: string[], salesList: number[]): void {
    for (let i = 0; i < gameNames.length; i++) {
      let sale: Sales = {
        name: gameNames[i],
        y: salesList[i],
        color: this.getRandomColor()
      };

      // console.log(sale)
      this.saleSeries.push(sale);
    }
  }

  initializeChart(): void {
    this.chart = new Chart({
      chart: {
        type: 'pie',
        height: 325
      },
      title: {
        text: 'Sales by Category'
      },
      xAxis: {
        categories: this.titles 
      },
      yAxis: {
        title: {
          text: 'Revenue in %'
        }
      },
      series: [
        {
          type: 'pie',
          data: this.saleSeries 
        }
      ],
      credits: {
        enabled: false
      }
    });
  }
  
  ngOnInit(): void {
    this.retrieveSales()
  }

}




// chart = new Chart({
//   chart: {
//     type: 'pie',
//     height: 325
//   },
//   title: {
//     text: 'Category wise sales'
//   },
//   xAxis: {
//     categories: [
//       'Electronics',
//       'Groceries',
//       'Cosmetics',
//       'Clothes',
//       'Appliances',
//     ]
//   },
//   yAxis: {
//     title: {
//       text: 'Revenue in %'
//     }
//   },
//   series: [
//     {
//       type: 'pie',
//       data: [
//         {
//           name: 'Electronics',
//           y: 41.0,
//           color: '#044342',
//         },
//         {
//           name: 'Groceries',
//           y: 33.8,
//           color: '#7e0505',
//         },
//         {
//           name: 'Cosmetics',
//           y: 6.5,
//           color: '#ed9e20',
//         },
//         {
//           name: 'Clothes',
//           y: 15.2,
//           color: '#6920fb',
//         },
//         {
//           name: 'Appliances',
//           y: 3.5,
//           color: '#121212',
//         },
//       ]
//     }
//   ],
//   credits: {
//     enabled: false
//   }
// })
