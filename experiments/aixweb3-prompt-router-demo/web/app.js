const q=document.getElementById('q');
const ask=document.getElementById('ask');
const route=document.getElementById('route');
const answer=document.getElementById('answer');
const related=document.getElementById('related');

document.querySelectorAll('[data-q]').forEach(b=>b.onclick=()=>{q.value=b.dataset.q;run()});
ask.onclick=run;

async function run(){
  const question=q.value.trim();
  if(!question)return;
  route.textContent='正在识别问题类型并选择 Prompt 模板...';
  answer.textContent='正在调用 Hermes 模型接口生成回答...';
  related.textContent='等待推荐相关概念...';
  try{
    const res=await fetch('/api/ask',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({question})});
    const data=await res.json();
    render(data);
  }catch(e){
    route.textContent='请求失败';
    answer.textContent='请确认 Go 服务已启动，并检查浏览器控制台。';
    related.textContent='';
  }
}

function render(d){
  route.innerHTML=`<div class="routeGrid">
    <div><b>问题类型</b><br>${d.type}</div>
    <div><b>Prompt 模板</b><br>${d.templateName}</div>
    <div><b>选择原因</b><br>${d.routeReason}</div>
  </div>`;
  answer.textContent=d.answer || d.error || '无回答';
  related.innerHTML='<b>继续学习：</b> '+(d.relatedConcepts||[]).map(x=>`<button class="relatedBtn" data-next="${x}">${x}</button>`).join(' ');
  related.querySelectorAll('[data-next]').forEach(b=>b.onclick=()=>{q.value=`请用适合初学者的方式解释 ${b.dataset.next}，并说明它和 ${d.question} 的关系。`;run()});
}
