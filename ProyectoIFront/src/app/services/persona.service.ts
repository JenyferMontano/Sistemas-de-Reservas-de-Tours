import {Injectable} from "@angular/core"
import {server} from "./global"
import {HttpClient,HttpHeaders} from "@angular/common/http"
import { Observable } from "rxjs"

@Injectable({
    providedIn:'root'
})export class PersonaService{
    public url:string
    private accessToken:string
    constructor(private _http:HttpClient){
        this.url=server.url
        this.accessToken=""
    }
    getPersonas():Observable<any>{
       const headers=new HttpHeaders().set('Content-Type','application/json')       
        const options={
            headers
        }
        return this._http.get(this.url+'persona',options)
    }
 
}