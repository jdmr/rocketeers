(function(e){function t(t){for(var s,r,i=t[0],c=t[1],l=t[2],d=0,m=[];d<i.length;d++)r=i[d],a[r]&&m.push(a[r][0]),a[r]=0;for(s in c)Object.prototype.hasOwnProperty.call(c,s)&&(e[s]=c[s]);u&&u(t);while(m.length)m.shift()();return o.push.apply(o,l||[]),n()}function n(){for(var e,t=0;t<o.length;t++){for(var n=o[t],s=!0,i=1;i<n.length;i++){var c=n[i];0!==a[c]&&(s=!1)}s&&(o.splice(t--,1),e=r(r.s=n[0]))}return e}var s={},a={app:0},o=[];function r(t){if(s[t])return s[t].exports;var n=s[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,r),n.l=!0,n.exports}r.m=e,r.c=s,r.d=function(e,t,n){r.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},r.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},r.t=function(e,t){if(1&t&&(e=r(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(r.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var s in e)r.d(n,s,function(t){return e[t]}.bind(null,s));return n},r.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return r.d(t,"a",t),t},r.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},r.p="/pbe/";var i=window["webpackJsonp"]=window["webpackJsonp"]||[],c=i.push.bind(i);i.push=t,i=i.slice();for(var l=0;l<i.length;l++)t(i[l]);var u=c;o.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"0139":function(e,t,n){},"21e5":function(e,t,n){"use strict";var s=n("9581"),a=n.n(s);a.a},"23c3":function(e,t,n){},"36ad":function(e,t,n){"use strict";var s=n("0139"),a=n.n(s);a.a},"506b":function(e,t,n){"use strict";var s=n("23c3"),a=n.n(s);a.a},"56d7":function(e,t,n){"use strict";n.r(t);n("cadf"),n("551c");var s=n("2b0e"),a=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-app",[n("navigation"),n("v-container",{attrs:{id:"main"}},[n("router-view")],1)],1)},o=[],r=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("v-toolbar",{attrs:{id:"nav-toolbar",fixed:"",app:""}},[n("div",{attrs:{id:"nav-container"}},[n("v-container",{attrs:{fluid:""}},[n("v-layout",{attrs:{row:"","align-center":""}},[n("v-flex",{attrs:{xs11:"",md4:""}},[n("v-toolbar-title",{attrs:{id:"nav-title"}},[n("router-link",{attrs:{id:"nav-title-link",to:{name:"home"}}},[e._v("\n              Keene Rocketeers\n            ")])],1)],1),e.mobile?n("v-flex",{attrs:{xs2:""}},[n("v-menu",{attrs:{"nudge-bottom":40,bottom:"",right:""}},[n("v-btn",{attrs:{slot:"activator",icon:""},slot:"activator"},[n("v-icon",[e._v("menu")])],1),n("v-list",[n("v-list-tile",[n("v-list-tile-title",{on:{click:function(t){e.$router.push({name:"home"})}}},[e._v("PBE")])],1),n("v-list-tile",[n("v-list-tile-title",{on:{click:function(t){e.$router.push({name:"questions"})}}},[e._v("Questions")])],1)],1)],1)],1):e._e(),e.mobile?e._e():n("v-flex",{attrs:{md8:""}},[n("v-toolbar-items",{attrs:{id:"menu-items"}},[n("v-layout",{attrs:{row:"","justify-end":""}},[n("v-btn",{staticClass:"menu-item",attrs:{to:{name:"home"},flat:"",exact:""}},[e._v("PBE")]),n("v-btn",{staticClass:"menu-item",attrs:{to:{name:"questions"},flat:"",exact:""}},[e._v("Questions")])],1)],1)],1)],1)],1)],1)])},i=[],c={name:"Navigation",data:function(){return{windowWidth:0,mobile:!1}},watch:{windowWidth:function(e){this.mobile=e<=500}},mounted:function(){var e=this;this.windowWidth=window.innerWidth,this.$nextTick(function(){window.addEventListener("resize",function(){e.windowWidth=window.innerWidth})})}},l=c,u=(n("61fe"),n("506b"),n("2877")),d=Object(u["a"])(l,r,i,!1,null,"4dc8d146",null),m=d.exports,v={name:"App",components:{Navigation:m}},f=v,h=(n("7ed5"),Object(u["a"])(f,a,o,!1,null,null,null)),g=h.exports,p=(n("7f7f"),n("8c4f")),b=n("323e"),_=n.n(b),k=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"main"},[n("h1",{staticClass:"display-2"},[e._v("Games")]),n("v-btn",{attrs:{large:"",color:"primary"},on:{click:function(t){e.addingGame=!0}}},[n("v-icon",[e._v("add")]),e._v(" New Game")],1),n("div",{staticClass:"games"},e._l(e.games,function(t){return n("v-card",{key:t.id},[n("v-card-title",{staticClass:"headline secondary",attrs:{"primary-title":""}},[n("span",{staticStyle:{color:"#ffffff"}},[e._v(e._s(t.name))])]),n("v-card-text",[n("v-list",[n("v-list-tile",[n("v-list-tile-content",[n("v-list-tile-title",[e._v(e._s(t.seconds))]),n("v-list-tile-sub-title",[e._v("Seconds")])],1)],1),n("v-list-tile",[n("v-list-tile-content",[n("v-list-tile-title",[e._v(e._s(e._f("date")(t.created)))]),n("v-list-tile-sub-title",[e._v("Created")])],1)],1),n("v-list-tile",[n("v-list-tile-content",[n("v-list-tile-title",[e._v(e._s(t.questions))]),n("v-list-tile-sub-title",[e._v("Questions")])],1)],1)],1),n("p",{staticClass:"title"},[e._v("Teams")]),n("v-list",{staticClass:"teams-div"},e._l(t.teams,function(s){return n("v-list-tile",{key:s.id},[n("v-list-tile-title",[e._v(e._s(s.name))]),n("v-list-tile-action",{staticClass:"team-join-action"},[n("v-btn",{attrs:{flat:"",icon:"",color:"secondary"},on:{click:function(n){e.openTeam(t,s)}}},[n("v-icon",[e._v("launch")])],1)],1)],1)}))],1),n("v-card-actions",[n("v-btn",{attrs:{color:"primary",large:"",block:""},on:{click:function(n){e.addTeam(t)}}},[n("v-icon",[e._v("launch")]),e._v(" Join")],1),"OPEN"===t.status?n("v-btn",{attrs:{color:"success",large:"",block:""},on:{click:function(n){e.startGame(t)}}},[n("v-icon",[e._v("play_circle_outline")]),e._v(" Start")],1):e._e(),"STARTED"===t.status?n("v-btn",{attrs:{color:"error",large:"",block:""},on:{click:function(n){e.finishGame(t)}}},[n("v-icon",[e._v("stop")]),e._v(" Finish")],1):e._e()],1),n("v-card-actions",[n("v-btn",{attrs:{color:"error",large:"",block:""},on:{click:function(n){e.deleteGame(t)}}},[n("v-icon",[e._v("delete")]),e._v(" Delete")],1)],1)],1)})),n("v-dialog",{attrs:{width:"400"},model:{value:e.addingGame,callback:function(t){e.addingGame=t},expression:"addingGame"}},[n("v-card",[n("v-card-title",{staticClass:"headline primary",attrs:{"primary-title":""}},[n("span",{staticStyle:{color:"#ffffff"}},[e._v("Add Game")])]),n("v-card-text",[n("v-text-field",{attrs:{label:"Game Name",required:""},model:{value:e.game.name,callback:function(t){e.$set(e.game,"name",t)},expression:"game.name"}}),n("v-text-field",{attrs:{label:"Question Seconds",required:"",type:"number",min:"1"},model:{value:e.game.seconds,callback:function(t){e.$set(e.game,"seconds",t)},expression:"game.seconds"}}),n("v-text-field",{attrs:{label:"Number of Questions",required:"",type:"number",min:"1"},model:{value:e.game.questions,callback:function(t){e.$set(e.game,"questions",t)},expression:"game.questions"}}),n("v-text-field",{attrs:{label:"Game Book",required:""},model:{value:e.book,callback:function(t){e.book=t},expression:"book"}}),n("v-text-field",{attrs:{label:"Game Chapters",required:"",hint:"Chapters separated by commas (E.g. 2,3,4,5,21)","persistent-hint":""},model:{value:e.chapters,callback:function(t){e.chapters=t},expression:"chapters"}})],1),n("v-divider"),n("v-card-actions",[n("v-spacer"),n("v-btn",{attrs:{disabled:e.cannotAddGame,color:"secondary"},on:{click:e.addGame}},[n("v-icon",[e._v("add")]),e._v(" Add\n        ")],1)],1)],1)],1),n("v-dialog",{attrs:{width:"400"},model:{value:e.addingTeam,callback:function(t){e.addingTeam=t},expression:"addingTeam"}},[n("v-card",[n("v-card-title",{staticClass:"headline primary",attrs:{"primary-title":""}},[n("span",{staticStyle:{color:"#ffffff"}},[e._v("Join Game")])]),n("v-card-text",[n("v-text-field",{attrs:{label:"Team Name",required:""},model:{value:e.team.name,callback:function(t){e.$set(e.team,"name",t)},expression:"team.name"}})],1),n("v-divider"),n("v-card-actions",[n("v-spacer"),n("v-btn",{attrs:{disabled:e.joining,color:"secondary"},on:{click:e.join}},[n("v-icon",[e._v("add")]),e._v(" Add\n        ")],1)],1)],1)],1),n("v-snackbar",{attrs:{timeout:5e3,color:"error"},model:{value:e.hasErrors,callback:function(t){e.hasErrors=t},expression:"hasErrors"}},[e._v("\n    "+e._s(e.errors)+"\n    "),n("v-btn",{attrs:{dark:"",flat:""},on:{click:function(t){e.hasErrors=!1}}},[e._v("\n      Close\n    ")])],1)],1)},w=[],C=(n("a481"),n("ac6a"),n("28a5"),n("bc3a")),q=n.n(C),y=q.a.create({});y.interceptors.request.use(function(e){return _.a.start(),e},function(e){return _.a.done(),Promise.reject(e)}),y.interceptors.response.use(function(e){return _.a.done(),e},function(e){return _.a.done(),Promise.reject(e)});var x=y,D=n("70f2"),E=n.n(D),I={name:"Home",filters:{date:function(e){return E()(e,"dddd, MMM Do, YYYY h:mm:ss a")}},data:function(){return{wsUrl:"ws://keenechurch.org:9000",games:[],gettingGames:!1,errors:null,hasErrors:!1,addingGame:!1,game:{name:null,seconds:30,questions:10},creatingGame:!1,joining:!1,addingTeam:!1,team:{name:null},book:null,chapters:null,teamsWS:null,checkConnectionsInterval:null}},computed:{cannotAddGame:function(){return!this.game.name||!!this.creatingGame}},beforeDestroy:function(){clearInterval(this.checkConnectionsInterval)},created:function(){this.getGames(),this.connectTeamsWS();var e=this;e.checkConnectionsInterval=setInterval(function(){e.checkConnections()},5e3)},methods:{getGames:function(){var e=this;this.gettingGames=!0,x.get("/api/v1/games").then(function(t){e.games=t.data}).catch(function(t){console.error("Could not get games: ",t),e.errors="Could not get games: "+t.response.data,e.hasErrors=!0}).finally(function(){e.gettingGames=!1})},addGame:function(){var e=this;this.creatingGame=!0,this.game.chapters=[];var t=this.chapters.split(",");t.map(function(t){e.game.chapters.push({book:e.book,chapter:t})}),x.post("/api/v1/games",this.game).then(function(){e.getGames(),e.addingGame=!1}).catch(function(t){console.error("Could not add game: ",t.response),e.errors="Could not add game: "+t.response.data,e.hasErrors=!0}).finally(function(){e.creatingGame=!1})},addTeam:function(e){this.game=e,this.team={name:null},this.addingTeam=!0},join:function(){var e=this;this.joining=!0,x.post("/api/v1/games/"+this.game.id+"/teams",this.team).then(function(){e.getGame(e.game.id),e.addingTeam=!1}).catch(function(t){console.error("Could not join game: "+e.game.name,t.response),e.errors="Could not join game: "+e.game.name+" : "+t.response.data,e.hasErrors=!0}).finally(function(){e.joining=!1})},getGame:function(e){var t=this;x.get("/api/v1/games/"+e).then(function(n){t.game=n.data,t.games.forEach(function(n,s){n.id===e&&(t.games[s]=t.game)})}).catch(function(n){console.error("Could not get game: "+e,n.response),t.errors="Could not get game: "+e+" : "+n.response.data,t.hasErrors=!0})},startGame:function(e){var t=this;console.log("starting game: ",e),x.post("/api/v1/games/"+e.id+"/start").then(function(){t.$router.push({name:"game",params:{gameID:e.id}})}).catch(function(n){console.error("Could not get game started: ",e,n.response),t.errors="Could not get started: "+n.response.data,t.hasErrors=!0})},finishGame:function(e){var t=this;console.log("finishing game: ",e),x.post("/api/v1/games/"+e.id+"/finish").then(function(){t.teamsWS.send("finished game"),t.$router.replace({name:e,params:{gameID:e.id}})}).catch(function(n){console.error("Could not finish game: ",e,n.response),t.errors="Could not finish game: "+n.response.data,t.hasErrors=!0})},connectTeamsWS:function(){var e=this;this.teamsWS=new WebSocket(this.wsUrl+"/ws/pbe/teams"),this.teamsWS.onopen=function(){console.log("Connected to teams websocket"),e.teamsWS.send("connected")},this.teamsWS.onmessage=function(t){console.log("Received a team websocket update"),e.games=JSON.parse(t.data)}},checkConnections:function(){console.log("Checking connections"),this.teamsWS.readyState===this.teamsWS.CLOSED&&this.connectTeamsWS()},openTeam:function(e,t){this.$router.push({name:"team",params:{gameID:e.id,teamID:t.id}})},deleteGame:function(e){var t=this;confirm("Are you sure you want to delete game: "+e.name+"?")&&x.delete("/api/v1/games/"+e.id).then(function(){t.getGames(),t.teamsWS.send("deleted game")}).catch(function(n){console.error("Could not delete game: "+e.id,n.response),t.errors="Could not delete game: "+e.name+" : "+n.response.data,t.hasErrors=!0})}}},S=I,G=(n("36ad"),Object(u["a"])(S,k,w,!1,null,"130adec8",null)),W=G.exports,Q=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"main"},[n("h1",{staticClass:"display-2"},[e._v("Questions")]),n("v-btn",{attrs:{large:"",color:"primary"},on:{click:e.openAddQuestionDialog}},[n("v-icon",[e._v("add")]),e._v(" New Question")],1),n("div",{staticClass:"questions"},e._l(e.questions,function(t){return n("v-card",{key:t.id},[n("v-card-title",{staticClass:"headline secondary",attrs:{"primary-title":""}},[n("span",{staticStyle:{color:"#ffffff"}},[e._v(e._s(t.question))])]),n("v-card-text",[n("v-list",[n("v-list-tile",[n("v-list-tile-content",[n("v-list-tile-title",[e._v(e._s(t.book))]),n("v-list-tile-sub-title",[e._v("Book")])],1),n("v-list-tile-content",[n("v-list-tile-title",[e._v(e._s(t.chapter))]),n("v-list-tile-sub-title",[e._v("Chapter")])],1),n("v-list-tile-content",[n("v-list-tile-title",[e._v(e._s(t.verses))]),n("v-list-tile-sub-title",[e._v("Verses")])],1)],1)],1),n("v-list",e._l(t.answers,function(s){return n("v-list-tile",{key:s.id,attrs:{avatar:""}},[n("v-list-tile-avatar",[s.status?n("v-icon",{attrs:{color:"success"}},[e._v("check_circle")]):n("v-icon",{attrs:{color:"error"}},[e._v("block")])],1),n("v-list-tile-content",[n("v-list-tile-title",[e._v(e._s(s.answer))])],1),n("v-list-tile-action",[n("v-btn",{attrs:{icon:"",ripple:"",color:"error"},on:{click:function(n){e.deleteAnswer(t,s)}}},[n("v-icon",[e._v("delete")])],1)],1)],1)}))],1),n("v-card-actions",[n("v-btn",{attrs:{color:"primary",large:"",block:""},on:{click:function(n){e.openAddAnswerDialog(t)}}},[n("v-icon",[e._v("add")]),e._v(" Answer")],1),n("v-btn",{attrs:{color:"error",large:"",block:""},on:{click:function(n){e.deleteQuestion(t)}}},[n("v-icon",[e._v("delete")]),e._v(" Delete")],1)],1)],1)})),n("v-dialog",{attrs:{width:"400"},model:{value:e.addingQuestion,callback:function(t){e.addingQuestion=t},expression:"addingQuestion"}},[n("v-card",[n("v-card-title",{staticClass:"headline primary",attrs:{"primary-title":""}},[n("span",{staticStyle:{color:"#ffffff"}},[e._v("Add Question")])]),n("v-card-text",[n("v-textarea",{attrs:{label:"Question",required:""},model:{value:e.question.question,callback:function(t){e.$set(e.question,"question",t)},expression:"question.question"}}),n("v-text-field",{attrs:{label:"Book",required:""},model:{value:e.question.book,callback:function(t){e.$set(e.question,"book",t)},expression:"question.book"}}),n("v-text-field",{attrs:{label:"Chapter",required:""},model:{value:e.question.chapter,callback:function(t){e.$set(e.question,"chapter",t)},expression:"question.chapter"}}),n("v-text-field",{attrs:{label:"Verses",required:""},model:{value:e.question.verses,callback:function(t){e.$set(e.question,"verses",t)},expression:"question.verses"}})],1),n("v-divider"),n("v-card-actions",[n("v-spacer"),n("v-btn",{attrs:{color:"secondary"},on:{click:e.addQuestion}},[n("v-icon",[e._v("add")]),e._v(" Add\n        ")],1)],1)],1)],1),n("v-dialog",{attrs:{width:"400"},model:{value:e.addingAnswer,callback:function(t){e.addingAnswer=t},expression:"addingAnswer"}},[n("v-card",[n("v-card-title",{staticClass:"headline primary",attrs:{"primary-title":""}},[n("span",{staticStyle:{color:"#ffffff"}},[e._v("Add Answer")])]),n("v-card-text",[n("v-textarea",{attrs:{label:"Answer",required:""},model:{value:e.answer.answer,callback:function(t){e.$set(e.answer,"answer",t)},expression:"answer.answer"}}),n("v-switch",{attrs:{label:"Correct answer?",color:"primary"},model:{value:e.answer.status,callback:function(t){e.$set(e.answer,"status",t)},expression:"answer.status"}})],1),n("v-divider"),n("v-card-actions",[n("v-spacer"),n("v-btn",{attrs:{color:"secondary"},on:{click:e.addAnswer}},[n("v-icon",[e._v("add")]),e._v(" Add\n        ")],1)],1)],1)],1),n("v-snackbar",{attrs:{timeout:5e3,color:"error"},model:{value:e.hasErrors,callback:function(t){e.hasErrors=t},expression:"hasErrors"}},[e._v("\n    "+e._s(e.errors)+"\n    "),n("v-btn",{attrs:{dark:"",flat:""},on:{click:function(t){e.hasErrors=!1}}},[e._v("\n      Close\n    ")])],1)],1)},$=[],A={name:"Questions",data:function(){return{addingQuestion:!1,question:{question:null},addingAnswer:!1,answer:{answer:null,status:!1},hasErrors:!1,errors:null,questions:[]}},created:function(){this.getQuestions()},methods:{openAddQuestionDialog:function(){this.question={question:null},this.addingQuestion=!0},addQuestion:function(){var e=this;this.addingQuestion=!0,x.post("/api/v1/questions",this.question).then(function(){e.getQuestions(),e.addingQuestion=!1}).catch(function(t){console.error("Could not add question: ",t.response),e.errors="Could not add question: "+t.response.data,e.hasErrors=!0})},getQuestions:function(){var e=this;x.get("/api/v1/questions").then(function(t){e.questions=t.data}).catch(function(t){console.error("Could not get questions: ",t.response),e.errors="Could not get questions: "+t.response.data,e.hasErrors=!0})},getQuestion:function(e){var t=this;x.get("/api/v1/questions/"+e).then(function(n){t.question=n.data,t.questions.forEach(function(n,s){n.id===e&&(t.questions[s]=t.question)})}).catch(function(n){console.error("Could not get question: "+e,n.response),t.errors="Could not get question: "+e+" : "+n.response.data,t.hasErrors=!0})},openAddAnswerDialog:function(e){this.question=e,this.answer={answer:null,status:!1},this.addingAnswer=!0},addAnswer:function(){var e=this;x.post("/api/v1/questions/"+this.question.id+"/answers",this.answer).then(function(){e.getQuestion(e.question.id),e.addingAnswer=!1}).catch(function(t){console.error("Could not add answer: ",t.response),e.errors="Could not add answer: "+t.response.data,e.hasErrors=!0})},deleteQuestion:function(e){var t=this;x.delete("/api/v1/questions/"+e.id).then(function(){t.getQuestions()}).catch(function(e){console.error("Could not delete question: ",e.response),t.errors="Could not delete question: "+e.response.data,t.hasErrors=!0})},deleteAnswer:function(e,t){var n=this;x.delete("/api/v1/questions/"+e.id+"/answers/"+t.id).then(function(){n.getQuestion(e.id)}).catch(function(e){console.error("Could not delete answer: ",e.response),n.errors="Could not delete answer: "+e.response.data,n.hasErrors=!0})}}},j=A,T=(n("9e39"),Object(u["a"])(j,Q,$,!1,null,"562ad7b8",null)),O=T.exports,N=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"main"},[n("v-card",[n("v-card-title",{staticClass:"headline secondary",attrs:{"primary-title":""}},[n("span",{staticStyle:{color:"#ffffff"}},[e._v(e._s(e.question.question))])]),n("v-card-text",[n("v-list",e._l(e.question.answers,function(t){return n("v-list-tile",{key:t.id},[n("v-list-tile-action",[n("v-btn",{attrs:{icon:"",ripple:""},on:{click:function(n){e.answerQuestion(t)}}},[t.checked?n("v-icon",[e._v("check_box")]):n("v-icon",[e._v("check_box_outline_blank")])],1)],1),n("v-list-tile-content",[e._v(e._s(t.answer))])],1)}))],1),n("v-card-actions",[n("v-btn",{attrs:{color:"primary"},on:{click:e.previousQuestion}},[n("v-icon",[e._v("navigate_before")]),e._v(" Previous")],1),n("v-btn",{attrs:{color:"primary"},on:{click:e.nextQuestion}},[n("v-icon",[e._v("navigate_next")]),e._v(" Next")],1),n("v-btn",{attrs:{color:"error"},on:{click:e.finishGame}},[n("v-icon",[e._v("stop")]),e._v(" Finish")],1)],1)],1),n("v-snackbar",{attrs:{timeout:5e3,color:"error"},model:{value:e.hasErrors,callback:function(t){e.hasErrors=t},expression:"hasErrors"}},[e._v("\n    "+e._s(e.errors)+"\n    "),n("v-btn",{attrs:{dark:"",flat:""},on:{click:function(t){e.hasErrors=!1}}},[e._v("\n      Close\n    ")])],1)],1)},P=[],M=(n("7514"),{name:"Game",data:function(){return{wsUrl:"ws://keenechurch.org:9000",gameID:this.$route.params.gameID,question:{answers:[],question:null},errors:null,hasErrors:!1,gameWS:null,checkConnectionsInterval:null,team:{}}},beforeDestroy:function(){clearInterval(this.checkConnectionsInterval)},created:function(){this.getHomeTeam(),this.getCurrentQuestion(),this.connectGameWS();var e=this;e.checkConnectionsInterval=setInterval(function(){e.checkConnections()},5e3)},methods:{getHomeTeam:function(){var e=this;x.get("/api/v1/games/"+this.gameID+"/home").then(function(t){console.log("Got home team: ",t.data),e.team=t.data,e.question&&e.team.answers&&e.team.answers.map(function(t){var n=e.question.answers.find(function(e){return e.id===t.id});n&&(n.checked=!0)})}).catch(function(t){console.error("Could not get home team: ",t.response),e.errors="Could not get home team: "+t.response.data,e.hasErrors=!0})},getCurrentQuestion:function(){var e=this;x.get("/api/v1/games/"+this.gameID+"/current").then(function(t){e.question=t.data,e.team&&e.team.answers&&e.team.answers.map(function(t){var n=e.question.answers.find(function(e){return e.id===t.id});n&&(n.checked=!0)})}).catch(function(t){console.error("Could not get current question: ",t.response),e.errors="Could not get current question: "+t.response.data,e.hasErrors=!0})},nextQuestion:function(){var e=this;x.post("/api/v1/games/"+this.gameID+"/next").then(function(){e.gameWS.send(e.gameID)}).catch(function(t){console.error("Could not get next question: ",t.response),e.errors="Could not get next question: "+t.response.data,e.hasErrors=!0})},previousQuestion:function(){var e=this;x.post("/api/v1/games/"+this.gameID+"/previous").then(function(){e.gameWS.send(e.gameID)}).catch(function(t){console.error("Could not get previous question: ",t.response),e.errors="Could not get previous question: "+t.response.data,e.hasErrors=!0})},finishGame:function(){var e=this;x.post("/api/v1/games/"+this.gameID+"/finish").then(function(){e.gameWS.send(e.gameID),e.$router.replace({name:"finished",params:{gameID:e.gameID}})}).catch(function(t){console.error("Could not finish game: ",t.response),e.errors="Could not finish game: "+t.response.data,e.hasErrors=!0})},connectGameWS:function(){var e=this;this.gameWS=new WebSocket(this.wsUrl+"/ws/pbe/game/"+this.gameID),this.gameWS.onopen=function(){console.log("Connected to game websocket"),e.gameWS.send(e.gameID)},this.gameWS.onmessage=function(t){e.question=JSON.parse(t.data),e.question.finished?e.$router.replace({name:"finished",params:{gameID:e.gameID}}):e.getHomeTeam()}},checkConnections:function(){console.log("Checking connections"),this.gameWS.readyState===this.gameWS.CLOSED&&this.connectGameWS()},answerQuestion:function(e){var t=this;e.checked?(e.checked=!1,x.delete("/api/v1/games/"+this.gameID+"/teams/"+this.team.id+"/answers/"+e.id).catch(function(n){console.error("Could not delete answer: ",e.id,n.response),t.errors="Could not delete answer: "+n.response.data,t.hasErrors=!0})):(e.checked=!0,x.post("/api/v1/games/"+this.gameID+"/teams/"+this.team.id+"/answers/"+e.id).catch(function(n){console.error("Could not submit answer: ",e.id,n.response),t.errors="Could not submit answer: "+n.response.data,t.hasErrors=!0}))}}}),Y=M,F=Object(u["a"])(Y,N,P,!1,null,null,null),J=F.exports,B=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"main"},[n("v-card",[n("v-card-title",{staticClass:"headline secondary",attrs:{"primary-title":""}},[n("span",{staticStyle:{color:"#ffffff"}},[e._v(e._s(e.question.question))])]),n("v-card-text",[n("v-list",e._l(e.question.answers,function(t){return n("v-list-tile",{key:t.id},[n("v-list-tile-action",[n("v-btn",{attrs:{icon:"",ripple:""},on:{click:function(n){e.answerQuestion(t)}}},[t.checked?n("v-icon",[e._v("check_box")]):n("v-icon",[e._v("check_box_outline_blank")])],1)],1),n("v-list-tile-content",[e._v(e._s(t.answer))])],1)}))],1)],1),n("v-snackbar",{attrs:{timeout:5e3,color:"error"},model:{value:e.hasErrors,callback:function(t){e.hasErrors=t},expression:"hasErrors"}},[e._v("\n    "+e._s(e.errors)+"\n    "),n("v-btn",{attrs:{dark:"",flat:""},on:{click:function(t){e.hasErrors=!1}}},[e._v("\n      Close\n    ")])],1)],1)},U=[],H={name:"Game",data:function(){return{wsUrl:"ws://keenechurch.org:9000",gameID:this.$route.params.gameID,teamID:this.$route.params.teamID,team:{},question:{answers:[],question:null},errors:null,hasErrors:!1,gameWS:null,checkConnectionsInterval:null}},beforeDestroy:function(){clearInterval(this.checkConnectionsInterval)},created:function(){this.getTeam(),this.getCurrentQuestion(),this.connectGameWS();var e=this;e.checkConnectionsInterval=setInterval(function(){e.checkConnections()},5e3)},methods:{getTeam:function(){var e=this;x.get("/api/v1/games/"+this.gameID+"/teams/"+this.teamID).then(function(t){e.team=t.data,e.question&&e.team.answers&&e.team.answers.map(function(t){var n=e.question.answers.find(function(e){return e.id===t.id});n&&(n.checked=!0)})}).catch(function(t){console.error("Could not get team: "+e.teamID,t.response),e.errors="Could not get team: "+e.teamID+" : "+t.response.data,e.hasErrors=!0})},getCurrentQuestion:function(){var e=this;x.get("/api/v1/games/"+this.gameID+"/current").then(function(t){e.question=t.data,e.team&&e.team.answers&&e.team.answers.map(function(t){var n=e.question.answers.find(function(e){return e.id===t.id});n&&(n.checked=!0)})}).catch(function(t){console.error("Could not get current question: ",t.response),e.errors="Could not get current question: "+t.response.data,e.hasErrors=!0})},nextQuestion:function(){var e=this;x.post("/api/v1/games/"+this.gameID+"/next").then(function(){e.gameWS.send(e.gameID)}).catch(function(t){console.error("Could not get next question: ",t.response),e.errors="Could not get next question: "+t.response.data,e.hasErrors=!0})},previousQuestion:function(){var e=this;x.post("/api/v1/games/"+this.gameID+"/previous").then(function(){e.gameWS.send(e.gameID)}).catch(function(t){console.error("Could not get previous question: ",t.response),e.errors="Could not get previous question: "+t.response.data,e.hasErrors=!0})},finishGame:function(){var e=this;x.post("/api/v1/games/"+this.gameID+"/finish").then(function(){e.gameWS.send(e.gameID),e.$router.replace({name:"finished",params:{gameID:e.gameID}})}).catch(function(t){console.error("Could not finish game: ",t.response),e.errors="Could not finish game: "+t.response.data,e.hasErrors=!0})},connectGameWS:function(){var e=this;this.gameWS=new WebSocket(this.wsUrl+"/ws/pbe/game/"+this.gameID),this.gameWS.onopen=function(){console.log("Connected to game websocket"),e.gameWS.send(e.gameID)},this.gameWS.onmessage=function(t){e.question=JSON.parse(t.data),e.question.finished?e.$router.replace({name:"finished",params:{gameID:e.gameID}}):e.getTeam()}},checkConnections:function(){console.log("Checking connections"),this.gameWS.readyState===this.gameWS.CLOSED&&this.connectGameWS()},answerQuestion:function(e){var t=this;e.checked?(e.checked=!1,x.delete("/api/v1/games/"+this.gameID+"/teams/"+this.teamID+"/answers/"+e.id).catch(function(n){console.error("Could not delete answer: ",e.id,n.response),t.errors="Could not delete answer: "+n.response.data,t.hasErrors=!0})):(e.checked=!0,x.post("/api/v1/games/"+this.gameID+"/teams/"+this.teamID+"/answers/"+e.id).catch(function(n){console.error("Could not submit answer: ",e.id,n.response),t.errors="Could not submit answer: "+n.response.data,t.hasErrors=!0}))}}},L=H,R=Object(u["a"])(L,B,U,!1,null,null,null),V=R.exports,z=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"main"},[n("h1",{staticClass:"display-2"},[e._v("Finished Game")]),n("v-card",[n("v-card-title",{staticClass:"headline secondary",attrs:{"primary-title":""}},[n("span",{staticStyle:{color:"#ffffff"}},[e._v(e._s(e.game.name))])]),n("v-card-text",[n("v-list",[n("v-list-tile",[n("v-list-tile-content",[n("v-list-tile-title",[e._v(e._s(e.game.seconds))]),n("v-list-tile-sub-title",[e._v("Seconds")])],1)],1),n("v-list-tile",[n("v-list-tile-content",[n("v-list-tile-title",[e._v(e._s(e._f("date")(e.game.created)))]),n("v-list-tile-sub-title",[e._v("Created")])],1)],1),n("v-list-tile",[n("v-list-tile-content",[n("v-list-tile-title",[e._v(e._s(e.game.questions))]),n("v-list-tile-sub-title",[e._v("Questions")])],1)],1)],1),n("p",{staticClass:"title"},[e._v("Teams")]),n("v-list",{staticClass:"teams-div"},e._l(e.game.teams,function(t){return n("v-list-tile",{key:t.id},[n("v-list-tile-title",[e._v(e._s(t.name))]),n("v-list-tile-action",{staticClass:"team-join-action"},[e._v("\n            "+e._s(t.points)+"\n          ")])],1)}))],1)],1),n("v-snackbar",{attrs:{timeout:5e3,color:"error"},model:{value:e.hasErrors,callback:function(t){e.hasErrors=t},expression:"hasErrors"}},[e._v("\n    "+e._s(e.errors)+"\n    "),n("v-btn",{attrs:{dark:"",flat:""},on:{click:function(t){e.hasErrors=!1}}},[e._v("\n      Close\n    ")])],1)],1)},K=[],X={name:"Finished",filters:{date:function(e){return E()(e,"dddd, MMM Do, YYYY h:mm:ss a")}},data:function(){return{gameID:this.$route.params.gameID,game:{name:null,teams:[]},errors:null,hasErrors:!1}},created:function(){this.getGame()},methods:{getGame:function(){var e=this;x.get("/api/v1/games/"+this.gameID+"/finished").then(function(t){e.game=t.data}).catch(function(t){console.error("Could not get game: "+e.gameID,t.response),e.errors="Could not get game: "+e.gameID+" : "+t.response.data,e.hasErrors=!0})}}},Z=X,ee=(n("21e5"),Object(u["a"])(Z,z,K,!1,null,"9c093794",null)),te=ee.exports;s["default"].use(p["a"]);var ne=new p["a"]({mode:"history",base:"/pbe/",routes:[{path:"/",name:"home",component:W},{path:"/questions",name:"questions",component:O},{path:"/game/:gameID",name:"game",component:J},{path:"/game/:gameID/finished",name:"finished",component:te},{path:"/game/:gameID/team/:teamID",name:"team",component:V}],scrollBehavior:function(e,t,n){return n||{x:0,y:0}}});ne.beforeResolve(function(e,t,n){e.name&&_.a.start(),n()}),ne.afterEach(function(){_.a.done()});var se=ne,ae=n("ce5b"),oe=n.n(ae);s["default"].use(oe.a,{theme:{primary:"#2196F3",secondary:"#1976D2",accent:"#FF5252",error:"#f44336",warning:"#ffeb3b",info:"#2196f3",success:"#4caf50"}}),s["default"].config.productionTip=!1,new s["default"]({router:se,render:function(e){return e(g)}}).$mount("#app")},"61fe":function(e,t,n){"use strict";var s=n("e0da"),a=n.n(s);a.a},"7ed5":function(e,t,n){"use strict";var s=n("8eff"),a=n.n(s);a.a},"8eff":function(e,t,n){},9581:function(e,t,n){},"9e39":function(e,t,n){"use strict";var s=n("df83"),a=n.n(s);a.a},df83:function(e,t,n){},e0da:function(e,t,n){}});
//# sourceMappingURL=app.1e0bd37b.js.map