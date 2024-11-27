# Mandatory II

## The Assignments

Reflect on how your group uses version control. Inspired by [Branching Strategy](https://github.com/who-knows-inc/KEA_DAT_DevOps_2024_Autumn/blob/main/00._Course_Material/01._Assignments/04._Sofware_Quality_Linting_CI/02._After/choose_a_git_branching_strategy.md) and other Git/GitHub related assignments.

- [How are you DevOps?](https://github.com/who-knows-inc/KEA_DAT_DevOps_2024_Autumn/blob/main/00._Course_Material/01._Assignments/07._Guest_Lecture/02._After/how_are_you_devops.md)
- [Software Quality](https://github.com/who-knows-inc/KEA_DAT_DevOps_2024_Autumn/blob/main/00._Course_Material/01._Assignments/04._Sofware_Quality_Linting_CI/02._After/software_quality.md)
- [Monitoring Realization](https://github.com/who-knows-inc/KEA_DAT_DevOps_2024_Autumn/blob/main/00._Course_Material/01._Assignments/11._Searching_Logging_Monitoring/02._After/monitoring_realization.md)

## Refleksion af Version Control

I vores projekt har vi ikke været gode til at udgive releases, når der har været nok ændringer til at retfærdiggøre det.

Vi har tidligere sagt, at vi brugte en kombination af `Gitflow` & `GitHub Flow`, men efterhånden har vi skiftet til en kombination af `Feature Branching` & `GitHub Flow`. Vi opretter en feature branch ud fra hver issue og laver pull requests til `dev`. Vi bruger ikke GitHub Discussions, men anvender `Discord` til kommunikation blandt teammedlemmer.

### Trunk Based Branching

Vi valgte ikke `trunk-based branching`, da vi anså det for en udfordring snarere end en fordel. Vi har ikke erfaring med denne branching-strategi og mangler det nødvendige mindset. Derudover risikerer vi at glemme at køre `git fetch` før merger, hvilket er en af vores udfordringer i forhold til DevOps-principperne.

### Release Branching

Vi har ikke anvendt `Release Branching`, da vi ønskede at komme hurtigt i gang med det, vi allerede kendte til. Efter refleksion indser vi, at en `Release Branching`-strategi ville have givet os bedre kontrol over udgivelser og ændringer, samt muligheden for nem rollback ved problemer. Dette kunne have øget antallet af releases, hvis vi havde haft en klar plan for udgivelsesfrekvens.

### Gitflow

Oprindeligt troede vi, at vi brugte `Gitflow`, men i praksis anvendte vi `Feature Branching` og `GitHub Flow`. For at kunne hævde, at vi brugte noget lignende `Gitflow`, skulle vi have implementeret en `Release Branching`-strategi, hvilket krævede en mere detaljeret projektplan. Manglen på overordnet planlægning har gjort det klart, at vi ikke har arbejdet tilstrækkeligt i overensstemmelse med DevOps-principperne.

## How are you DevOps?

### Why

- Åben kommunikation i teamet.
- Fysiske møder og `pair programming` muliggør hurtig problemløsning.
- Opsætning af `continuous integration` og `continuous delivery` nærmer os DevOps-principperne.
- Vi har skiftet hosting-provider undervejs, og da dette skift skulle ske havde vi en glidende overgang hvor begge services var aktive, til at vi kunne verificere at det virkede på den nye hosting service.

### Why Not

- Manglende fokus på at implementere tests tidligt i udviklingsprocessen har hæmmet integrationen i vores CI-Pipeline.
- Ingen implementering af `continuous deployment`, hvilket forhindrer opnåelse af fuld DevOps.
- Manglende automatisering af releases har resulteret i nedetid og manuel upload af nye releases.

## Software Quality

Ifølge SonarQube samt Code Climate har vi opnået en acceptabel grad af Software Quality med en `Maintainability status` på Grade A.

- Ifølge SonarCloud har vi mindre end 5% teknisk gæld, hvilket bidrager til vores Grade A rating for Maintainability.
- Vi er enige med SonarCloud og CodeClimate metrics, da vi selv kan bekræfte, at der ikke er duplicate code i vores projekt.
- Vi forbedrede vores security rating ved at flytte `legacy koden` til et andet repository, så det ikke forstyrrer vores nye repository.

Selvfølgelig kan man ikke følge disse værktøjer slavisk og stole blindt på dem, da den opdager duplications ved f.eks samme kode der står ved `PUT` og `POST` request koden, derfor skal man tjekke det igennem selv og ikke acceptere alle kriterierne uden verificering.

```typescript
{
  method: "POST",
  headers: {
    "Content-Type": "application/json",
}

{
  method: "PUT",
  headers: {
    "Content-Type": "application/json",
}
```

## Monitoring Realization

Ved at sætte vores monitoring op, har vi fået et overblik over vores programs CPU- og memory-forbrug, som var markant lavere end forventet. Vores Azure-setup bruger ikke så meget kraft som tildelt, hvilket muliggør skalering ned til en VM med mindre memory og CPU, og dermed besparelse af omkostninger.

Desuden har monitoringen hjulpet os med at identificere og løse bugs relateret til databasen og vores endpoints, som vi ikke var opmærksomme på tidligere.
