# 🌟 Refleksion

## ⚙️ Valg af Framework
Vi har valgt **Gorilla** som vores framework. 
👉 Årsagen er, at det er et letvægts-framework, som er nemt at lære og bruge. Det gør det nemt, når man kun har brug for lette pakker til at bygge en webapplikation.

## 💻 Valg af Programmeringssprog
Vi har valgt **Go** som programmeringssprog til vores backend.
🔍 Go blev også valgt, fordi vi har set det blive brugt i produktion, og vi ønskede at lære et nyt sprog i stedet for at bruge JavaScript.

## 🗄️ Valg af Database
Vi har valgt **SQLite** som database, men har planer om at skifte til **PostgreSQL** i fremtiden.
👉 I begyndelsen brugte vi SQLite, da det var det, der blev givet fra det gamle projekt. Skiftet til PostgreSQL blev valgt, fordi vi så det som en god mulighed for at lære en ny database, men samtidig have noget genkendeligt, da vi tidligere har arbejdet med MySQL.

## 🌐 Valg af Frontend Framework
Vi startede med at bruge **Svelte** til vores frontend, men efter at have opdaget, at Gorilla havde server-side rendering, besluttede vi os for at prøve det. 
❗ Efter at have haft svært ved at forstå, hvordan det fungerer, valgte vi at vende tilbage til **Svelte** til frontend.

## ☁️ Valg af Deployments-tjeneste
Vi valgte at bruge **Azure** som vores deployment-tjeneste.
💡 Vi har brugt Azure før, og vi har set det anvendt i produktion. Det er nemt at bruge, gratis til open source-projekter, og vi har mulighed for at bruge det gratis som studerende med nogle studiekreditter.

## 📝 Valg af Pull Request-skabelon
Vi valgte at bruge en skabelon til pull requests.
✔️ Skabelonen gør det nemmere for os at forstå, hvad der kræves i en pull request og sikrer, at vi er konsekvente i resten af projektet. Det gør det også lettere for os at gennemgå og godkende pull requests.

## 📋 Valg af Issue-skabelon
Vi har valgt at bruge en issue-skabelon for at gøre det nemmere at forstå, hvad der skal gøres i et issue.
❓ Hvis nogen er usikre på, hvad der skal gøres, kan de bede om hjælp i issue-tråden.

## 🌳 Valg af Branching-strategi
Vi har valgt en slags hybrid af **Feature Branching** og **GitHub flow** som vores branching-strategi. 
🔀 Vi har vores `Prod`, `Dev` og `Testing` branches. 
- `Prod`-branchen er vores produktionsgren
- `Dev`-branchen er vores udviklingsgren
- `Testing`-branchen bruges til testning og merges ind i `Prod`-branchen, når den er klar til næste udgivelse.

📌 Derudover har vi også konceptet med feature branches og pull requests, som skal gennemgås, inden de merges ind i `Dev`-branchen. På denne måde sikrer vi, at koden er af høj kvalitet, og fejl fanges, inden de merges i `Dev`-branchen.

## 🚀 Valg af Continuous Integration og Continuous Delivery
Vi har valgt at bruge **GitHub Actions** til Continuous Integration og Continuous Delivery.
🛠️ Vi har brugt GitHub Actions før, men vi har ikke brugt det til kontinuerlig integration og levering før. Vi valgte GitHub Actions, fordi det er nemt at bruge og gratis til open source-projekter. Det var også et valg, vi traf for at spare tid på opsætningen af CI/CD-pipelinen.
