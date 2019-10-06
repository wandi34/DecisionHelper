import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DecisionComponent } from './decision/decision-stack/decision-stack.component';



const routes: Routes = [
  {path: '', component: DecisionComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
