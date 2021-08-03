import { Redirect, Route, Switch } from 'react-router-dom';
import NotFound from 'components/Content/NotFound';
import Pawn from 'components/Content/Pawns/Pawn';
import Import from 'components/Content/File/Import';
import Export from 'components/Content/File/Export';
import Replace from 'components/Content/Pawns/Replace';

const Content = () => (
  <main className="col-md-8 col-lg-9 col-xl-9 ml-sm-auto px-md-4">
    <div id="content">
      <Switch>
        <Route exact path="/">
          <Redirect to="/import-file" />
        </Route>
        <Route exact path="/import-file">
          <Import />
        </Route>
        <Route exact path="/export-file">
          <Export />
        </Route>
        <Route path="/pawns/:category(main|first|second)" exact component={Pawn} />
        <Route path="/pawns/:category(main|first|second)/replace" exact component={Replace} />
        <Route component={NotFound} />
      </Switch>
    </div>
  </main>
);

export default Content;
